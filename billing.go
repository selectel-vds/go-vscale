package vscale_api_go

import (
	"fmt"
	"net/http"
)

type BillingService struct {
	client Client
}

type Billing struct {
	Balance int64  `json:"balance,omitempty"`
	Bonus   int64  `json:"bonus,omitempty"`
	Status  string `json:"status,omitempty"`
	Summ    int64  `json:"summ,omitempty"`
	Unpaid  int64  `json:"unpaid,omitempty"`
	UserID  int64  `json:"user_id,omitempty"`
}

type Payments struct {
	Items []struct {
		Created int64  `json:"created,omitempty"`
		Desc    string `json:"desc,omitempty"`
		Dir     int64  `json:"dir,omitempty"`
		ID      int64  `json:"id,omitempty"`
		IsBonus int64  `json:"is_bonus,omitempty"`
		Price   int64  `json:"price,omitempty"`
		State   int64  `json:"state,omitempty"`
		Type    int64  `json:"type,omitempty"`
	} `json:"items,omitempty"`
	Status string `json:"status,omitempty"`
}

type Consumption struct {
	Huge struct {
		Count int64 `json:"count,omitempty"`
		Summ  int64 `json:"summ,omitempty"`
	} `json:"huge,omitempty"`
	Large struct {
		Count int64 `json:"count,omitempty"`
		Summ  int64 `json:"summ,omitempty"`
	} `json:"large,omitempty"`
	Medium struct {
		Count int64 `json:"count,omitempty"`
		Summ  int64 `json:"summ,omitempty"`
	} `json:"medium,omitempty"`
	Monster struct {
		Count int64 `json:"count,omitempty"`
		Summ  int64 `json:"summ,omitempty"`
	} `json:"moster,omitempty"`
	Small struct {
		Count int64 `json:"count,omitempty"`
		Summ  int64 `json:"summ,omitempty"`
	} `json:"small,omitempty"`
	Summ int64 `json:"summ,omitempty"`
}

func (c *BillingService) Billing() (*Billing, *http.Response, error) {

	billing := new(Billing)

	res, err := c.client.ExecuteRequest("GET", "billing/balance", []byte{}, billing)

	return billing, res, err
}

func (c *BillingService) Payments() (*Payments, *http.Response, error) {

	payments := new(Payments)

	res, err := c.client.ExecuteRequest("GET", "billing/payments/", []byte{}, payments)

	return payments, res, err
}

func (c *BillingService) Consumption(start, end string) (*map[string]Consumption, *http.Response, error) {

	consumption := make(map[string]Consumption)

	res, err := c.client.ExecuteRequest("GET", fmt.Sprintf("billing/consumption?start=%s&end=%s", start, end), []byte{}, &consumption)

	return &consumption, res, err
}
