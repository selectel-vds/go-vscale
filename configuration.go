package vscale_api_go

import (
	"net/http"
)

type ConfigurationService struct {
	client Client
}

type Rplan struct {
	Addresses int64    `json:"addresses,omitempty"`
	CPUs      int64    `json:"cpus,omitempty"`
	Locations []string `json:"locations,omitempty"`
	ID        string   `json:"id,omitempty"`
	Memory    int64    `json:"memory,omitempty"`
	Templates []string `json:"templates,omitempty"`
	Network   int64    `json:"network,omitempty"`
	Disk      int64    `json:"disk,omitempty"`
}

type Prices struct {
	Default struct {
		Backup int64 `json:"backup,omitempty"`
		Huge   struct {
			Hour  int64 `json:"hour,omitempty"`
			Month int64 `json:"month,omitempty"`
		} `json:"huge,omitempty"`
		Large struct {
			Hour  int64 `json:"hour,omitempty"`
			Month int64 `json:"month,omitempty"`
		} `json:"large,omitempty"`
		Medium struct {
			Hour  int64 `json:"hour,omitempty"`
			Month int64 `json:"month,omitempty"`
		} `json:"medium,omitempty"`
		Monster struct {
			Hour  int64 `json:"hour,omitempty"`
			Month int64 `json:"month,omitempty"`
		} `json:"monster,omitempty"`
		Network int64 `json:"network,omitempty"`
		Small   struct {
			Hour  int64 `json:"hour,omitempty"`
			Month int64 `json:"month,omitempty"`
		} `json:"small,omitempty"`
	} `json:"default,omitempty"`
	Period string `json:"period,omitempty"`
}

func (c *ConfigurationService) ListRplans() (*[]Rplan, *http.Response, error) {

	rplans := new([]Rplan)

	res, err := c.client.ExecuteRequest("GET", "rplans", []byte{}, rplans)

	return rplans, res, err
}

func (c *ConfigurationService) ListPrices() (*Prices, *http.Response, error) {

	prices := new(Prices)

	res, err := c.client.ExecuteRequest("GET", "billing/prices", []byte{}, prices)

	return prices, res, err
}
