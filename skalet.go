package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type SkaletService struct {
	client Client
}

type Skalet struct {
	Hostname        string        `json:"hostname,omitempty"`
	Locked          bool          `json:"locked,omitempty"`
	Locations       string        `json:"locations,omitempty"`
	Rplan           string        `json:"rplan,omitempty"`
	Name            string        `json:"name,omitempty"`
	Active          bool          `json:"active,omitempty"`
	Keys            []Keys        `json:"keys,omitempty"`
	PublicAddresses PublicAddress `json:"public_address,omitempty"`
	Status          string        `json:"status,omitempty"`
	MadeFrom        string        `json:"made_from,omitempty"`
	CTID            int64         `json:"ctid,omitempty"`
}

type Keys struct {
	Name string `json:"name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

type PublicAddress struct {
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Address string `json:"address,omitempty"`
}

type Task struct {
	Location string `json:"location,omitempty"`
	DInsert  string `json:"d_insert,omitempty"`
	ID       string `json:"id,omitempty"`
	Done     bool   `json:"done,omitempty"`
	Scalet   int64  `json:"scalet,omitempty"`
	Error    bool   `json:"error,omitempty"`
	DStart   string `json:"d_start,omitempty"`
	Method   string `json:"method,omitempty"`
	DEnd     string `json:"d_end,omitempty"`
}

func (s *SkaletService) List() (*[]Skalet, *http.Response, error) {

	skalets := new([]Skalet)

	res, err := s.client.ExecuteRequest("GET", "scalets", []byte{}, skalets)

	return skalets, res, err
}

func (s *SkaletService) Create(makeFrom, rplan, name, password, location string, doStart bool,
	keys []int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	body := struct {
		MakeFrom string  `json:"make_from,omitempty"`
		Rplan    string  `json:"rplan,omitempty"`
		DoStart  bool    `json:"do_start,omitempty"`
		Name     string  `json:"name,omitempty"`
		Keys     []int64 `json:"keys,omitempty"`
		Password string  `json:"password,omitempty"`
		Location string  `json:"location,omitempty"`
	}{makeFrom, rplan, doStart, name, keys, password, location}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("POST", "scalets", b, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Remove(CTID int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("DELETE", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Get(CTID int64) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	res, err := s.client.ExecuteRequest("GET", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, skalet)

	return skalet, res, err
}

func (s *SkaletService) Restart(CTID int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/restart"), []byte{}, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Rebuild(CTID int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/rebuild"), []byte{}, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Stop(CTID int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/stop"), []byte{}, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Start(CTID int64, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/start"), []byte{}, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Upgrade(CTID int64, rplan string, wait bool) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	body := struct {
		Rplan string `json:"rplan,omitempty"`
	}{rplan}

	b, _ := json.Marshal(body)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("POST", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/upgrade"), b, skalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return skalet, res, err
	}

	return skalet, res, err
}

func (s *SkaletService) Tasks() (*[]Task, *http.Response, error) {

	tasks := new([]Task)

	res, err := s.client.ExecuteRequest("GET", "tasks", []byte{}, tasks)

	return tasks, res, err
}

func (s *SkaletService) AddSSHKeys(CTID int64, keys []int64) (*Skalet, *http.Response, error) {

	skalet := new(Skalet)

	body := struct {
		Keys []int64 `json:"keys,omitempty"`
	}{keys}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), b, skalet)

	return skalet, res, err
}
