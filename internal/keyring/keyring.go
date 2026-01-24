package keyring

import (
	"github.com/zalando/go-keyring"
)

func SaveTokens(accessToken, refreshToken string) error {
	err := keyring.Set("sympal", "access_token", accessToken)
	if err != nil {
		return err
	}
	err = keyring.Set("sympal", "refresh_token", refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func LoadTokens() (accessToken, refreshToken string, err error) {
	accessToken, err = keyring.Get("sympal", "access_token")
	if err != nil {
		return "", "", err
	}

	refreshToken, err = keyring.Get("sympal", "refresh_token")
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
