package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"

	"github.com/david-fitzgerald/sympal/internal/keyring"
)

const (
	redirectURL = "http://localhost:8080/callback"
	authURL     = "https://accounts.google.com/o/oauth2/auth"
	tokenURL    = "https://oauth2.googleapis.com/token"
	scope       = "https://www.googleapis.com/auth/calendar.readonly"
)

type authResult struct {
	code  string
	error error
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func exchangeToken(code, clientID, clientSecret string) (*tokenResponse, error) {
	resp, err := http.PostForm(tokenURL, url.Values{"code": {code}, "client_id": {clientID}, "client_secret": {clientSecret}, "redirect_uri": {redirectURL}, "grant_type": {"authorization_code"}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token tokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func NewToken(refreshtoken, clientID, clientSecret string) (*tokenResponse, error) {
	resp, err := http.PostForm(tokenURL, url.Values{"grant_type": {"refresh_token"}, "refresh_token": {refreshtoken}, "client_id": {clientID}, "client_secret": {clientSecret}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token tokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func Authenticate(clientID, clientSecret string) (*tokenResponse, error) {
	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate state: %w", err)
	}
	resultChan := make(chan authResult)
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", makeCallbackHandler(state, resultChan))
	server := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		server.ListenAndServe()
	}()

	fullURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=code&state=%s",
		authURL,
		clientID,
		redirectURL,
		scope,
		state,
	)

	fmt.Println("Opening browser for authentication...")
	if err := openBrowser(fullURL); err != nil {
		return nil, fmt.Errorf("Failed to open browser: %w", err)
	}

	result := <-resultChan
	server.Shutdown(context.Background())

	if result.error != nil {
		return nil, result.error
	}

	token, err := exchangeToken(result.code, clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	err = keyring.SaveTokens(token.AccessToken, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func generateState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b) // crypto/rand, not math/rand
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func openBrowser(url string) error {
	return exec.Command("open", url).Start()
}

func makeCallbackHandler(expectedState string, resultChan chan authResult) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		code := query.Get("code")
		state := query.Get("state")

		if state != expectedState {
			resultChan <- authResult{error: fmt.Errorf("state mismatch")}
			w.Write([]byte("Error: state mismatch"))
			return
		}

		if code == "" {
			resultChan <- authResult{error: fmt.Errorf("code empty")}
			w.Write([]byte("Error: code empty"))
			return
		}

		resultChan <- authResult{code: code}
		w.Write([]byte("Success! You can close this tab."))

	}
}
