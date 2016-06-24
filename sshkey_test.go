package vscale_api_go

import (
	"testing"
)

// Test data
var Key string = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4jpT+KSXtUCWerNHTnqQd9qmyDEow84lKdp12/R/gfM28BQ+KbMMVJQmrWSNf1zEMmMTsiNR2h/20iw2V/guxRzF4jJm5NRwodKCwsZrAhjbGe0yMKG7GpHe50DhF8enxAOudcEMCBekrFlUrT+nd4bJVZ6ChrBN4BDaDlpKbJ9JN7lENB4Bs357K7hCahT3PPa+w/GnQ30vip5bi7BZEoOW2sgnF3BJS8IYPTNs253581PmvpepPpTM8TnE+tjvHHt6lwIbHw2nj7Diihik0r/TrBGJMFACkEh+MupIw0jzSh0GAhnZHJ8QuCCnVfScC04Fz1031QRzNPwK0rh4t bogdankurnosov@MacBook-Pro-Bogdan.local`
var KeyName string = "test_key"
var TestKeyID int64 = 4679

func TestSSHKeyService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.SSHKey.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestSSHKeyService_Create(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	sshkey, _, err := client.SSHKey.Create(Key, KeyName)
	if err != nil {
		t.Error(err)
		return
	}

	if sshkey.Name == "" {
		t.Error("SSH key's name is empty")
	}

	return

}

func TestSSHKeyService_Remove(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.SSHKey.Remove(TestKeyID)
	if err != nil {
		t.Error(err)
		return
	}

	return

}
