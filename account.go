package vscale_api_go

import (
	"net/http"
)

type AccountService struct {
	client Client
}

type Account struct {
	Info struct {
		Actdate    string `json:"actdate,omitempty"`
		Country    string `json:"country,omitempty"`
		Email      string `json:"email,omitempty"`
		FaceID     string `json:"face_id,omitempty"`
		ID         string `json:"id,omitempty"`
		Locale     string `json:"locale,omitempty"`
		Middlename string `json:"middlename,omitempty"`
		Mobile     string `json:"mobile,omitempty"`
		Name       string `json:"name,omitempty"`
		State      string `json:"state,omitempty"`
		Surname    string `json:"surname,omitempty"`
	} `json:"info,omitempty"`
}

func (a *AccountService) Get() (*Account, *http.Response, error) {

	account := new(Account)

	res, err := a.client.ExecuteRequest("GET", "account", []byte{}, account)

	return account, res, err
}
