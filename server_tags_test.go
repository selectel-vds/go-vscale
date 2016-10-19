package vscale_api_go

import (
	"testing"
)

var ServersTagName string = "vpn_scalets"
var ServersTagID int64 = 3

func TestServerTagService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.ServerTag.Create(ServersTagName, []int64{CTID})
	if err != nil {
		t.Error(err)
		return
	}

	if tag.Name == "" {
		t.Error("Tag name can't be empty")
		return
	}

	return
}

func TestServerTagService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerTag.List()
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func TestServerTagService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.ServerTag.Get(ServersTagID)
	if err != nil {
		t.Error(err)
		return
	}

	if tag.Name == "" {
		t.Error("Tag name can't be empty")
		return
	}

	return

}

func TestServerTagService_Update(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.ServerTag.Update(ServersTagID, "new_name", []int64{49174})
	if err != nil {
		t.Error(err)
		return
	}

	if tag.Name == "" {
		t.Error("Tag name can't be empty")
		return
	}

	return

}

func TestServerTagService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.ServerTag.Remove(ServersTagID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}
