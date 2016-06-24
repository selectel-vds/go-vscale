package vscale_api_go

import (
	"testing"
)

var PTRRecordContent string = "example.com"
var PTRRecordIP string = "95.213.195.160"
var PTRRecordID int64 = 2293

func TestPTRRecordService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	ptrRecord, _, err := client.PTRRecord.Create("123123", "1232")
	if err != nil {
		t.Error(err)
		return
	}

	if ptrRecord.ID == 0 {
		t.Error("PTR record id can't be empty")
	}

	return
}

func TestPTRRecordService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.PTRRecord.List()

	if err != nil {
		t.Error(err)
		return
	}

	return
}

func TestPTRRecordService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	record, _, err := client.PTRRecord.Get(PTRRecordID)

	if err != nil {
		t.Error(err)
		return
	}

	if record.ID == 0 {
		t.Error("Record ID can't be empty")
		return
	}

	return
}

func TestPTRRecordService_Update(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	ptrRecord, _, err := client.PTRRecord.Update(PTRRecordID, "example2.com", PTRRecordIP)

	if err != nil {
		t.Error(err)
		return
	}

	if ptrRecord.ID == 0 {
		t.Error("PTR record id can't be empty")
	}

	return
}

func TestPTRRecordService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.PTRRecord.Remove(PTRRecordID)

	if err != nil {
		t.Error(err)
		return
	}

	return
}
