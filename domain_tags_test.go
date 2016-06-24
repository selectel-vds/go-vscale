package vscale_api_go

import (
	"testing"
)

var DomainsTagName string = "group_2"
var DomainsTagID int64 = 3

func TestDomainTagService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.DomainTag.Create(DomainsTagName, []int64{TestDomainID})
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

func TestDomainTagService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.DomainTag.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestDomainTagService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.DomainTag.Get(DomainsTagID)
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

func TestDomainTagService_Update(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	tag, _, err := client.DomainTag.Update(DomainsTagID, "new_name", []int64{2541})
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

func TestDomainTagService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.DomainTag.Remove(DomainsTagID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}
