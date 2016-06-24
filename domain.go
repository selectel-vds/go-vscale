package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DomainService struct {
	client Client
}

type Domain struct {
	ID         int64       `json:"id,omitempty"`
	ChangeDate int64       `json:"change_date,omitempty"`
	CreateDate int64       `json:"create_date,omitempty"`
	Name       string      `json:"name,omitempty"`
	Tags       []DomainTag `json:"tags,omitempty"`
	UserID     int64       `json:"user_id,omitempty"`
}

func (d *DomainService) List() (*[]Domain, *http.Response, error) {

	domains := new([]Domain)

	res, err := d.client.ExecuteRequest("GET", "domains/", []byte{}, domains)

	return domains, res, err
}

func (d *DomainService) Create(name string) (*Domain, *http.Response, error) {

	// TODO Add bind zone

	domain := new(Domain)

	body := struct {
		Name string `json:"name,omitempty"`
	}{name}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("POST", "domains/", b, domain)

	return domain, res, err
}

func (d *DomainService) Get(domainID int64) (*Domain, *http.Response, error) {

	domain := new(Domain)

	res, err := d.client.ExecuteRequest("GET", fmt.Sprint("domains/", strconv.FormatInt(domainID, 10)), []byte{}, domain)

	return domain, res, err
}

func (d *DomainService) Update(domainID int64, tags []int64) (*Domain, *http.Response, error) {

	domain := new(Domain)

	body := struct {
		Tags []int64 `json:"tags,omitempty"`
	}{tags}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("PATCH", fmt.Sprint("domains/", strconv.FormatInt(domainID, 10)), b, domain)

	return domain, res, err
}

func (d *DomainService) Remove(domainID int64) (bool, *http.Response, error) {

	res, err := d.client.ExecuteRequest("DELETE", fmt.Sprint("domains/", strconv.FormatInt(domainID, 10)), []byte{}, struct{}{})

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
