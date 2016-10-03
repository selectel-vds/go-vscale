package vscale_api_go

import (
	"testing"
)

func TestBackupService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Backup.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestBackupService_Get(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Backup.Get(BackupID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestBackupService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Backup.Remove(BackupID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}
