package vscale_api_go

import (
	"testing"
)

var TestDomain string = "example.com"
var TestDomainID int64 = 2540
var DomainTags []int64 = []int64{543}

func TestDomainService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.DomainService.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestDomainService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	domain, _, err := client.DomainService.Create(TestDomain)
	if err != nil {
		t.Error(err)
		return
	}

	if domain.ID == 0 {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	domain, _, err := client.DomainService.Get(TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	if domain.ID == 0 {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Update(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	domain, _, err := client.DomainService.Update(TestDomainID, DomainTags)
	if err != nil {
		t.Error(err)
		return
	}

	if domain.ID == 0 {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	status, _, err := client.DomainService.Remove(TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	if !status {
		t.Fail()
		return
	}

	return

}
