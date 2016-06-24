package vscale_api_go

import (
	"testing"
)

// Test data
var CTID int64 = 35586
var KeyID int64 = 4639
var AdditionalKeyID int64 = 4674
var UpgradePlan string = "medium"

func TestServerService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerService.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestServerService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Create("ubuntu_16.04_64_001_master", "small", "test-1",
		"", "spb0", true, []int64{KeyID}, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name can\t be empty")
		return
	}

	return
}

func TestServerService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerService.Remove(35758, true)
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestServerService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Get(CTID)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Restart(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Restart(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Rebuild(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Rebuild(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Stop(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Stop(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Start(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Start(35758, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Uprade(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	server, _, err := client.ServerService.Upgrade(35758, UpgradePlan, true)

	if err != nil {
		t.Error(err)
		return
	}

	if server.Name == "" {
		t.Error("Server name is empty")
		return
	}

	return
}

func TestServerService_Tasks(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerService.Tasks()

	if err != nil {
		t.Error(err)
		return
	}

	return
}

func TestServerService_AddSSHKeys(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerService.AddSSHKeys(CTID, []int64{AdditionalKeyID})

	if err != nil {
		t.Error(err)
		return
	}

	return
}