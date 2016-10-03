package vscale_api_go

import (
	"fmt"
	"net/http"
)

type BackupService struct {
	client Client
}

func (s *BackupService) List() (*[]Backup, *http.Response, error) {

	backups := new([]Backup)

	res, err := s.client.ExecuteRequest("GET", "backups", []byte{}, backups)

	return backups, res, err
}

func (s *BackupService) Get(BackupID string) (*Backup, *http.Response, error) {

	backup := new(Backup)

	res, err := s.client.ExecuteRequest("GET", fmt.Sprintf("backups/%s", BackupID), []byte{}, backup)

	return backup, res, err
}

func (s *BackupService) Remove(BackupID string) (*Backup, *http.Response, error) {

	backup := new(Backup)

	res, err := s.client.ExecuteRequest("DELETE", fmt.Sprintf("backups/%s", BackupID), []byte{}, backup)

	return backup, res, err
}
