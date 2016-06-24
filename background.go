package vscale_api_go

import (
	"net/http"
)

type BackgroundService struct {
	client Client
}

type Location struct {
	Active            bool     `json:"active,omitempty"`
	PrivateNetworking bool     `json:"private_networking,omitempty"`
	Image             []string `json:"image,omitempty"`
	ID                string   `json:"id,omitempty"`
	Description       string   `json:"description,omitempty"`
	Rplans            []string `json:"rplans,omitempty"`
}

type Image struct {
	Rplans      []string `json:"rplans,omitempty"`
	Active      bool     `json:"active,omitempty"`
	Size        int64    `json:"size,omitempty"`
	Locations   []string `json:"locations,omitempty"`
	ID          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
}

func (b *BackgroundService) ListLocations() (*[]Location, *http.Response, error) {

	dcs := new([]Location)

	res, err := b.client.ExecuteRequest("GET", "locations", []byte{}, dcs)

	return dcs, res, err
}

func (b *BackgroundService) ListImages() (*[]Image, *http.Response, error) {

	images := new([]Image)

	res, err := b.client.ExecuteRequest("GET", "images", []byte{}, images)

	return images, res, err
}
