package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DomainTagService struct {
	client Client
}

type DomainTag struct {
	Name    string  `json:"name,omitempty"`
	Domains []int64 `json:"domains,omitempty"`
	ID      int64   `json:"id,omitempty"`
}

type DomainTagInfo struct {
	Name    string   `json:"name,omitempty"`
	Domains []Domain `json:"domains,omitempty"`
	ID      int64    `json:"id,omitempty"`
}

func (d *DomainTagService) Create(name string, domains []int64) (*DomainTag, *http.Response, error) {

	domainTag := new(DomainTag)

	body := struct {
		Name    string  `json:"name,omitempty"`
		Domains []int64 `json:"domains,omitempty"`
	}{name, domains}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("POST", "domains/tags/", b, domainTag)

	return domainTag, res, err
}

func (d *DomainTagService) List() (*[]DomainTag, *http.Response, error) {

	domainTags := new([]DomainTag)

	res, err := d.client.ExecuteRequest("GET", "domains/tags/", []byte{}, domainTags)

	return domainTags, res, err
}

func (d *DomainTagService) Get(tagID int64) (*DomainTagInfo, *http.Response, error) {

	domainTag := new(DomainTagInfo)

	url := fmt.Sprint("domains/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("GET", url, []byte{}, domainTag)

	return domainTag, res, err
}

func (d *DomainTagService) Update(tagID int64, name string, domains []int64) (*DomainTag, *http.Response, error) {

	domainTag := new(DomainTag)

	body := struct {
		Name    string  `json:"name,omitempty"`
		Domains []int64 `json:"domains,omitempty"`
	}{name, domains}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, domainTag)

	return domainTag, res, err
}

func (d *DomainTagService) Remove(tagID int64) (bool, *http.Response, error) {

	url := fmt.Sprint("domains/tags/", strconv.FormatInt(tagID, 10))

	res, err := d.client.ExecuteRequest("DELETE", url, []byte{}, struct{}{})

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
