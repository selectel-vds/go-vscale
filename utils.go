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

func IsSuccess(code int) bool {
	
	successCodes := []int{200, 201, 202, 203, 204, 205, 206, 207}
	
	for _, i := range successCodes {
		if i == code {
			return true
		}
	}
	
	return false
}
