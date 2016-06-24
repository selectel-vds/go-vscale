package vscale_api_go

import (
	"testing"
)

// Test data
var CTID int64 = 35586
var KeyID int64 = 4639
var AdditionalKeyID int64 = 4674
var UpgradePlan string = "medium"

func TestScaletService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Scalet.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestScaletService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Create("ubuntu_16.04_64_001_master", "small", "test-1",
		"", "spb0", true, []int64{KeyID}, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name can\t be empty")
		return
	}

	return
}

func TestScaletService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Scalet.Remove(35758, true)
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestScaletService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Get(CTID)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Restart(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Restart(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Rebuild(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Rebuild(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Stop(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Stop(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Start(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Start(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Uprade(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	scalet, _, err := client.Scalet.Upgrade(35758, UpgradePlan, true)

	if err != nil {
		t.Error(err)
		return
	}

	if scalet.Name == "" {
		t.Error("Scalet name is empty")
		return
	}

	return
}

func TestScaletService_Tasks(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Scalet.Tasks()

	if err != nil {
		t.Error(err)
		return
	}

	return
}

func TestScaletService_AddSSHKeys(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Scalet.AddSSHKeys(CTID, []int64{AdditionalKeyID})

	if err != nil {
		t.Error(err)
		return
	}

	return
}
