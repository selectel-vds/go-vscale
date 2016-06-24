package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DomainRecordService struct {
	client Client
}

type DomainRecord struct {
	TTL     int64  `json:"ttl,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

func (d *DomainRecordService) List(domainID int64) (*[]DomainRecord, *http.Response, error) {

	domainRecords := new([]DomainRecord)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/")

	res, err := d.client.ExecuteRequest("GET", url, []byte{}, domainRecords)

	return domainRecords, res, err
}

func (d *DomainRecordService) Create(domainID int64, name, recordType string, ttl int64, content string) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		TTL     int64  `json:"ttl,omitempty"`
		Content string `json:"content,omitempty"`
	}{name, recordType, ttl, content}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/")

	res, err := d.client.ExecuteRequest("POST", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) CreateSOA(domainID int64, name string, ttl int64, content, email string) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		TTL     int64  `json:"ttl,omitempty"`
		Content string `json:"content,omitempty"`
		Email   string `json:"email,omitempty"`
	}{name, "SOA", ttl, content, email}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/")

	res, err := d.client.ExecuteRequest("POST", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) CreateSRV(domainID int64, name string, ttl, weight, port int64,
	target string, priority int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name     string `json:"name,omitempty"`
		Type     string `json:"type,omitempty"`
		TTL      int64  `json:"ttl,omitempty"`
		Weight   int64  `json:"weight,omitempty"`
		Port     int64  `json:"port,omitempty"`
		Target   string `json:"target,omitempty"`
		Priority int64  `json:"priority,omitempty"`
	}{name, "SRV", ttl, weight, port, target, priority}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/")

	res, err := d.client.ExecuteRequest("POST", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) CreateMX(domainID int64, name string, ttl int64, content string, priority int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name     string `json:"name,omitempty"`
		Type     string `json:"type,omitempty"`
		TTL      int64  `json:"ttl,omitempty"`
		Content  string `json:"content,omitempty"`
		Priority int64  `json:"priority,omitempty"`
	}{name, "MX", ttl, content, priority}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/")

	res, err := d.client.ExecuteRequest("POST", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) Update(domainID, recordID int64, name, recordType string, ttl int64, content string) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		TTL     int64  `json:"ttl,omitempty"`
		Content string `json:"content,omitempty"`
	}{name, recordType, ttl, content}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) UpdateSOA(domainID, recordID int64, name string, ttl int64, content, email string) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		TTL     int64  `json:"ttl,omitempty"`
		Content string `json:"content,omitempty"`
		Email   string `json:"email,omitempty"`
	}{name, "SOA", ttl, content, email}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) UpdateSRV(domainID, recordID int64, name string, ttl, weight, port int64,
	target string, priority int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name     string `json:"name,omitempty"`
		Type     string `json:"type,omitempty"`
		TTL      int64  `json:"ttl,omitempty"`
		Weight   int64  `json:"weight,omitempty"`
		Port     int64  `json:"port,omitempty"`
		Target   string `json:"target,omitempty"`
		Priority int64  `json:"priority,omitempty"`
	}{name, "SRV", ttl, weight, port, target, priority}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) UpdateMX(domainID, recordID int64, name string, ttl int64, content string,
	priority int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	body := struct {
		Name     string `json:"name,omitempty"`
		Type     string `json:"type,omitempty"`
		TTL      int64  `json:"ttl,omitempty"`
		Content  string `json:"content,omitempty"`
		Priority int64  `json:"priority,omitempty"`
	}{name, "MX", ttl, content, priority}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("PUT", url, b, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) Get(domainID, recordID int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("GET", url, []byte{}, domainRecord)

	return domainRecord, res, err
}

func (d *DomainRecordService) Remove(domainID, recordID int64) (*DomainRecord, *http.Response, error) {

	domainRecord := new(DomainRecord)

	url := fmt.Sprint("domains/", strconv.FormatInt(domainID, 10), "/records/", strconv.FormatInt(recordID, 10))

	res, err := d.client.ExecuteRequest("DELETE", url, []byte{}, domainRecord)

	return domainRecord, res, err
}
