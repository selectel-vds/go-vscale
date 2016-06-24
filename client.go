package vscale_api_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
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
	AccountService       *AccountService
	ServerService        *ServerService
	BackgroundService    *BackgroundService
	ConfigurationService *ConfigurationService
	SSHKeyService        *SSHKeyService
	NotificationService  *NotificationService
	BillingService       *BillingService
	DomainService        *DomainService
	DomainRecordService  *DomainRecordService
	DomainTagService     *DomainTagService
	PTRRecordService     *PTRRecordService
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
	client.AccountService = &AccountService{client}
	client.ServerService = &ServerService{client}
	client.BackgroundService = &BackgroundService{client}
	client.ConfigurationService = &ConfigurationService{client}
	client.SSHKeyService = &SSHKeyService{client}
	client.NotificationService = &NotificationService{client}
	client.BillingService = &BillingService{client}
	client.DomainService = &DomainService{client}
	client.DomainRecordService = &DomainRecordService{client}
	client.DomainTagService = &DomainTagService{client}
	client.PTRRecordService = &PTRRecordService{client}

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

	if res.StatusCode != 200 && res.StatusCode != 204 {
		data, _ := ioutil.ReadAll(reader)

		APIerror := new(Error)

		json.Unmarshal(data, &APIerror)

		if APIerror.Error != "" {
			if APIerror.Field != "" {
				return res, errors.New(fmt.Sprint(APIerror.Field, ": ", APIerror.Error))
			}
			return res, errors.New(APIerror.Error)
		}

		return res, errors.New(string(data))
	}

	if object != nil {
		err := json.NewDecoder(reader).Decode(object)

		// EOF means empty response body, this error is not needed
		if err != nil && err != io.EOF {
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
