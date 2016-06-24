package vscale_api_go

import (
	"testing"
)

func TestAccountService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	account, _, err := client.Account.Get()
	if err != nil {
		t.Error(err)
		return
	}

	if account.Info.Name == "" {
		t.Error("Account name is empty")
		return
	}

	return

}
