package vscale_api_go

import (
	"errors"
	"os"
)

func GetToken() (string, error) {

	token := os.Getenv("VSCALE_API_TOKEN")
	if token == "" {
		return token, errors.New("Token is empty")
	}

	return token, nil
}
