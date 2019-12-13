package vscale_api_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Interface of Client for mocking data receiving in tests
type Client interface {
	ExecuteRequest(method, url string, body []byte, object interface{}) (*http.Response, error)
	WSSConn() (*websocket.Conn, error)
	WaitTask(c *websocket.Conn, taskID string) (bool, error)
}

type WebClient struct {
	Token      string
	HTTPClient *http.Client

	// TODO Change to type url
	// In fact it is host url and API version, example: https://api.vscale.io/v1/
	BaseURL string

	// Services which is used for accessing API
	Account       *AccountService
	Scalet        *ScaletService
	Background    *BackgroundService
	Configuration *ConfigurationService
	SSHKey        *SSHKeyService
	Notification  *NotificationService
	Billing       *BillingService
	Domain        *DomainService
	DomainRecord  *DomainRecordService
	DomainTag     *DomainTagService
	PTRRecord     *PTRRecordService
	Backup        *BackupService
	ServerTag     *ServerTagService
}

type Error struct {
	Error string `json:"error,omitempty"`
	Field string `json:"field,omitempty"`
}

// Web client creating
func NewClient(token string) *WebClient {

	// TODO Maybe it will be better to check if token is empty

	client := &WebClient{
		Token:      token,
		HTTPClient: new(http.Client),
		BaseURL:    "https://api.vscale.io/v1/",
	}

	// TODO Maybe it will be better to add account checking here via token, to be sure that token is valid and user exists

	// Passing client to all services for easy client mocking in future and not passing it to every function
	client.Account = &AccountService{client}
	client.Scalet = &ScaletService{client}
	client.Background = &BackgroundService{client}
	client.Configuration = &ConfigurationService{client}
	client.SSHKey = &SSHKeyService{client}
	client.Notification = &NotificationService{client}
	client.Billing = &BillingService{client}
	client.Domain = &DomainService{client}
	client.DomainRecord = &DomainRecordService{client}
	client.DomainTag = &DomainTagService{client}
	client.PTRRecord = &PTRRecordService{client}
	client.Backup = &BackupService{client}
	client.ServerTag = &ServerTagService{client}

	return client

}

// Executing HTTP Request (receiving info from API)
func (client *WebClient) ExecuteRequest(method, url string, body []byte, object interface{}) (*http.Response, error) {

	req, err := http.NewRequest(method, fmt.Sprint(client.BaseURL, url), bytes.NewBuffer(body))
	if err != nil {
		return new(http.Response), err
	}

	req.Header.Set("X-Token", client.Token)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	// Cloning response body for future using
	buf, _ := ioutil.ReadAll(res.Body)
	reader := ioutil.NopCloser(bytes.NewBuffer(buf))

	res.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if res.Header.Get("Vscale-Error-Message") != "None" && res.Header.Get("Vscale-Error-Message") != "" {
		return res, errors.New(res.Header.Get("Vscale-Error-Message"))
	}
	
	if !IsSuccess(res.StatusCode) {
		return res, errors.New("Not successful status code")
	}

	if object != nil && (res.StatusCode == 200 || res.StatusCode == 201) {
		err := json.NewDecoder(reader).Decode(object)

		// EOF means empty response body, this error is not needed
		if err != nil && err != io.EOF {
			log.Println(err)
			return res, err
		}
	}

	return res, nil
}

func (client *WebClient) WSSConn() (*websocket.Conn, error) {
	u := url.URL{Scheme: "wss", Host: "ws.api.vscale.io", Path: "/v1/ws"}
	header := http.Header{}
	header["X-Token"] = []string{client.Token}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return c, err
	}
	return c, nil
}

// Waiting until operation will be complete
func (client *WebClient) WaitTask(c *websocket.Conn, taskID string) (bool, error) {

	defer c.Close()

	msg := struct {
		ID     int64 `json:"id,omitempty"`
		Result struct {
			Type    string `json:"type,omitempty"`
			Error   string `json:"error,omitempty"`
			Name    string `json:"name,omitempty"`
			Status  string `json:"status,omitempty"`
			Message string `json:"message,omitempty"`
			Done    bool   `json:"done,omitempty"`
			ID      string `json:"id,omitempty"`
		} `json:"result,omitempty"`
		Time string `json:"time,omitempty"`
	}{}

	for {
		code, message, err := c.ReadMessage()
		if err != nil {
			return false, err
		}

		if code == 1 {
			json.Unmarshal(message, &msg)

			if msg.Result.ID == taskID && msg.Result.Done == true {
				if msg.Result.Error != "" {
					return false, errors.New(msg.Result.Error)
				}
				return true, nil
			}

		}

	}
}
