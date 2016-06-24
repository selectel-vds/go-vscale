package vscale_api_go

import (
	"encoding/json"
	"net/http"
)

type NotificationService struct {
	client Client
}

type Notification struct {
	NotifyBalance int64  `json:"notify_balance,omitempty"`
	Status        string `json:"status,omitempty"`
}

func (n *NotificationService) BillingSettings() (*Notification, *http.Response, error) {

	notification := new(Notification)

	res, err := n.client.ExecuteRequest("GET", "billing/notify", []byte{}, notification)

	return notification, res, err
}

func (n *NotificationService) BillingSettingsUpdate(notifyBalance int64) (*Notification, *http.Response, error) {

	// TODO Doesn't work because API don't use JSON in this case

	notification := new(Notification)

	body := struct {
		NotifyBalance int64 `json:"notify_balance,omitempty"`
	}{notifyBalance}

	b, _ := json.Marshal(body)

	res, err := n.client.ExecuteRequest("PUT", "billing/notify", b, notification)

	return notification, res, err
}
