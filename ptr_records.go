package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PTRRecordService struct {
	client Client
}

type PTRRecord struct {
	Content string `json:"content,omitempty"`
	UserID  int64  `json:"user_id,omitempty"`
	IP      string `json:"ip,omitempty"`
	ID      int64  `json:"id,omitempty"`
}

func (p *PTRRecordService) Create(content, ip string) (*PTRRecord, *http.Response, error) {

	ptrRecord := new(PTRRecord)

	body := struct {
		Content string `json:"content,omitempty"`
		IP      string `json:"ip,omitempty"`
	}{content, ip}

	b, _ := json.Marshal(body)

	res, err := p.client.ExecuteRequest("POST", "domains/ptr/", b, ptrRecord)

	return ptrRecord, res, err
}

func (p *PTRRecordService) List() (*[]PTRRecord, *http.Response, error) {

	ptrRecords := new([]PTRRecord)

	res, err := p.client.ExecuteRequest("GET", "domains/ptr/", []byte{}, ptrRecords)

	return ptrRecords, res, err
}

func (p *PTRRecordService) Get(recordID int64) (*PTRRecord, *http.Response, error) {

	ptrRecord := new(PTRRecord)

	url := fmt.Sprint("domains/ptr/", strconv.FormatInt(recordID, 10))

	res, err := p.client.ExecuteRequest("GET", url, []byte{}, ptrRecord)

	return ptrRecord, res, err
}

func (p *PTRRecordService) Update(recordID int64, content, ip string) (*PTRRecord, *http.Response, error) {

	ptrRecord := new(PTRRecord)

	body := struct {
		Content string `json:"content,omitempty"`
		IP      string `json:"ip,omitempty"`
	}{content, ip}

	b, _ := json.Marshal(body)

	url := fmt.Sprint("domains/ptr/", strconv.FormatInt(recordID, 10))

	res, err := p.client.ExecuteRequest("PUT", url, b, ptrRecord)

	return ptrRecord, res, err
}

func (p *PTRRecordService) Remove(recordID int64) (bool, *http.Response, error) {

	url := fmt.Sprint("domains/ptr/", strconv.FormatInt(recordID, 10))

	res, err := p.client.ExecuteRequest("DELETE", url, []byte{}, struct{}{})

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
