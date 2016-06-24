package vscale_api_go

import (
	"testing"
)

func TestConfigurationService_ListRplans(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	rplans, _, err := client.ConfigurationService.ListRplans()
	if err != nil {
		t.Error(err)
		return
	}

	if len(*rplans) == 0 {
		t.Error("No plans received")
		return
	}

	return

}

func TestConfigurationService_ListPrices(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ConfigurationService.ListPrices()
	if err != nil {
		t.Error(err)
		return
	}

	return

}
