package vscale_api_go

import (
	"testing"
)

// Test data
var CTID int64 = 35586
var KeyID int64 = 4639
var AdditionalKeyID int64 = 4674
var UpgradePlan string = "medium"

func TestSkaletService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Skalet.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestSkaletService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Create("ubuntu_16.04_64_001_master", "small", "test-1",
		"", "spb0", true, []int64{KeyID}, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name can\t be empty")
		return
	}

	return
}

func TestSkaletService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Skalet.Remove(35758, true)
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestSkaletService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Get(CTID)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Restart(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Restart(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Rebuild(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Rebuild(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Stop(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Stop(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Start(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Start(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Uprade(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	skalet, _, err := client.Skalet.Upgrade(35758, UpgradePlan, true)

	if err != nil {
		t.Error(err)
		return
	}

	if skalet.Name == "" {
		t.Error("Skalet name is empty")
		return
	}

	return
}

func TestSkaletService_Tasks(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Skalet.Tasks()

	if err != nil {
		t.Error(err)
		return
	}

	return
}

func TestSkaletService_AddSSHKeys(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Skalet.AddSSHKeys(CTID, []int64{AdditionalKeyID})

	if err != nil {
		t.Error(err)
		return
	}

	return
}
