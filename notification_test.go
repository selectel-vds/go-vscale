package vscale_api_go

import (
	"io/ioutil"
	"log"
	"testing"
)

// Test data
var BillingLimit int64 = 10000

func TestNotificationService_BillingSettings(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.NotificationService.BillingSettings()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestNotificationService_BillingSettingsUpdate(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, res, err := client.NotificationService.BillingSettingsUpdate(BillingLimit)
	data, _ := ioutil.ReadAll(res.Body)
	log.Println(string(data))
	if err != nil {
		t.Error(err)
		return
	}

	return

}
