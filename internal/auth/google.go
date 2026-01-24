package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os/exec"
)

const (
	redirectURL = "http://localhost:8080/callback"
	authURL     = "https://accounts.google.com/o/oauth2/auth"
	tokenURL    = "https://oauth2.googleapis.com/token"
	scope       = "https://www.googleapis.com/auth/calendar.readonly"
)

func Authenticate(clientID, clientSecret string) error {
	state, err := generateState()
	if err != nil {
		return fmt.Errorf("Failed to generate state: %w", err)
	}

	fullURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=code&state=%s",
		authURL,
		clientID,
		redirectURL,
		scope,
		state,
	)

	fmt.Println("Opening browser for authentication...")
	if err := openBrowser(fullURL); err != nil {
		return fmt.Errorf("Failed to open browser: %w", err)
	}

	return nil
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
