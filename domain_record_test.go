package vscale_api_go

import (
	"testing"
)

var DomainRecordName string = "example.com"
var DomainRecordID int64 = 13644
var DomainRecordContent string = "3.1.3.37"
var SOAEmail string = "gofort@example.com"

func TestDomainRecordService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.DomainRecordService.List(TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestDomainRecordService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.Create(TestDomainID, DomainRecordName, "A", 300, "3.1.3.37")
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_CreateSOA(t *testing.T) {

	// TODO Doesn't work, reason: unknown, response: {"error": "cant_add_soa"}

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.CreateSOA(TestDomainID, DomainRecordName, 300, DomainRecordContent, SOAEmail)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_CreateSRV(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.CreateSRV(TestDomainID, DomainRecordName, 300, 3000, 3000, "example.com", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_CreateMX(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.CreateMX(TestDomainID, DomainRecordName, 300, "example.com", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_Update(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.Update(TestDomainID, DomainRecordID, DomainRecordName, "A", 300, "3.1.3.38")
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_UpdateSOA(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.UpdateSOA(TestDomainID, DomainRecordID, DomainRecordName, 300, DomainRecordContent, SOAEmail)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_UpdateSRV(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.UpdateSRV(TestDomainID, DomainRecordID, DomainRecordName, 300, 3000, 3000, "example.com", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_UpdateMX(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.UpdateMX(TestDomainID, DomainRecordID, DomainRecordName, 300, "example.com", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.DomainRecordService.Get(TestDomainID, DomainRecordID)
	if err != nil {
		t.Error(err)
		return
	}

	if record.Name == "" {
		t.Error("Record name can't be empty")
		return
	}

	return

}

func TestDomainRecordService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.DomainRecordService.Remove(TestDomainID, DomainRecordID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}
