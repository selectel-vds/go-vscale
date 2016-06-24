package vscale_api_go

import (
	"testing"
)

func TestBackgroundService_ListDataCenters(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	dcs, _, err := client.BackgroundService.ListLocations()
	if err != nil {
		t.Error(err)
		return
	}

	if len(*dcs) == 0 {
		t.Error("No datacenters received")
		return
	}

	return

}

func TestBackgroundService_ListImages(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	images, _, err := client.BackgroundService.ListImages()
	if err != nil {
		t.Error(err)
		return
	}

	if len(*images) == 0 {
		t.Error("No images received")
		return
	}

	return

}
