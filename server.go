package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type ServerService struct {
	client Client
}

type Server struct {
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

func (s *ServerService) List() (*[]Server, *http.Response, error) {

	servers := new([]Server)

	res, err := s.client.ExecuteRequest("GET", "scalets", []byte{}, servers)

	return servers, res, err
}

func (s *ServerService) Create(makeFrom, rplan, name, password, location string, doStart bool,
	keys []int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

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

	res, err := s.client.ExecuteRequest("POST", "scalets", b, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Remove(CTID int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("DELETE", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Get(CTID int64) (*Server, *http.Response, error) {

	server := new(Server)

	res, err := s.client.ExecuteRequest("GET", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, server)

	return server, res, err
}

func (s *ServerService) Restart(CTID int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/restart"), []byte{}, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Rebuild(CTID int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/rebuild"), []byte{}, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Stop(CTID int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/stop"), []byte{}, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Start(CTID int64, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/start"), []byte{}, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Upgrade(CTID int64, rplan string, wait bool) (*Server, *http.Response, error) {

	server := new(Server)

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

	res, err := s.client.ExecuteRequest("POST", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/upgrade"), b, server)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return server, res, err
	}

	return server, res, err
}

func (s *ServerService) Tasks() (*[]Task, *http.Response, error) {

	tasks := new([]Task)

	res, err := s.client.ExecuteRequest("GET", "tasks", []byte{}, tasks)

	return tasks, res, err
}

func (s *ServerService) AddSSHKeys(CTID int64, keys []int64) (*Server, *http.Response, error) {

	server := new(Server)

	body := struct {
		Keys []int64 `json:"keys,omitempty"`
	}{keys}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), b, server)

	return server, res, err
}
