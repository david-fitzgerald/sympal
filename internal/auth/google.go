package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
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

func Authenticate(clientID, clientSecret string) (string, error) {
	state, err := generateState()
	if err != nil {
		return "", fmt.Errorf("Failed to generate state: %w", err)
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
		return "", fmt.Errorf("Failed to open browser: %w", err)
	}

	result := <-resultChan
	server.Shutdown(context.Background())

	if result.error != nil {
		return "", result.error
	}
	return result.code, nil
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
