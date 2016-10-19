package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ServerTagService struct {
	client Client
}

type ServerTag struct {
	Name    string  `json:"name,omitempty"`
	Scalets []int64 `json:"scalets,omitempty"`
	ID      int64   `json:"id,omitempty"`
}

func (d *ServerTagService) Create(name string, scalets []int64) (*ServerTag, *http.Response, error) {

	serverTag := new(ServerTag)

	body := struct {
		Name    string  `json:"name,omitempty"`
		Scalets []int64 `json:"scalets,omitempty"`
	}{name, scalets}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("POST", "scalets/tags", b, serverTag)

	return serverTag, res, err
}

func (d *ServerTagService) List() (*[]ServerTag, *http.Response, error) {

	serverTags := new([]ServerTag)

	res, err := d.client.ExecuteRequest("GET", "scalets/tags", []byte{}, serverTags)

	return serverTags, res, err
}

func (d *ServerTagService) Get(tagID int64) (*ServerTag, *http.Response, error) {

	serverTag := new(ServerTag)

	url := fmt.Sprint("scalets/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("GET", url, []byte{}, serverTag)

	return serverTag, res, err
}

func (d *ServerTagService) Update(tagID int64, name string, scalets []int64) (*ServerTag, *http.Response, error) {

	serverTag := new(ServerTag)

	body := struct {
		Name    string  `json:"name,omitempty"`
		Scalets []int64 `json:"scalets,omitempty"`
	}{name, scalets}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("scalets/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, serverTag)

	return serverTag, res, err
}

func (d *ServerTagService) Remove(tagID int64) (bool, *http.Response, error) {

	url := fmt.Sprint("scalets/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("DELETE", url, []byte{}, struct{}{})

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
