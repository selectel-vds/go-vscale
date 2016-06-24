package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SSHKeyService struct {
	client Client
}

type SSHKey struct {
	ID   int64  `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *SSHKeyService) List() (*[]SSHKey, *http.Response, error) {

	sshkeys := new([]SSHKey)

	res, err := a.client.ExecuteRequest("GET", "sshkeys", []byte{}, sshkeys)

	return sshkeys, res, err
}

func (s *SSHKeyService) Create(key, name string) (*SSHKey, *http.Response, error) {

	sshkey := new(SSHKey)

	body := struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
	}{key, name}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("POST", "sshkeys", b, sshkey)

	return sshkey, res, err
}

func (s *SSHKeyService) Remove(keyID int64) (bool, *http.Response, error) {

	res, err := s.client.ExecuteRequest("DELETE", fmt.Sprint("sshkeys/", strconv.FormatInt(keyID, 10)), []byte{}, struct{}{})

	if err != nil {
		return false, res, nil
	}

	return true, res, nil
}
