package vscale_api_go

import (
	"testing"
)

func TestBillingService_Billing(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	billing, _, err := client.Billing.Billing()
	if err != nil {
		t.Error(err)
		return
	}

	if billing.UserID == 0 {
		t.Error("Empty user id")
		return
	}

	return

}

func TestBillingService_Payments(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Billing.Payments()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestBillingService_Consumption(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Billing.Consumption("2016-06-01", "2016-06-30")
	if err != nil {
		t.Error(err)
		return
	}

	return

}
