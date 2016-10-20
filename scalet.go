package vscale_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type ScaletService struct {
	client Client
}

type Scalet struct {
	Hostname        string        `json:"hostname,omitempty"`
	Locked          bool          `json:"locked,omitempty"`
	Location        string        `json:"location,omitempty"`
	Rplan           string        `json:"rplan,omitempty"`
	Name            string        `json:"name,omitempty"`
	Active          bool          `json:"active,omitempty"`
	Keys            []Keys        `json:"keys,omitempty"`
	PublicAddresses PublicAddress `json:"public_address,omitempty"`
	Status          string        `json:"status,omitempty"`
	MadeFrom        string        `json:"made_from,omitempty"`
	CTID            int64         `json:"ctid,omitempty"`
}

type Backup struct {
	ID       string  `json:"id,omitempty"`
	Template string  `json:"template,omitempty"`
	Active   bool    `json:"active"`
	Name     string  `json:"name"`
	Scalet   int64   `json:"scalet"`
	Status   string  `json:"status"`
	Size     float64 `json:"size"`
	Locked   bool    `json:"locked"`
	Location string  `json:"location"`
	Created  string  `json:"created"`
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

func (s *ScaletService) List() (*[]Scalet, *http.Response, error) {

	scalets := new([]Scalet)

	res, err := s.client.ExecuteRequest("GET", "scalets", []byte{}, scalets)

	return scalets, res, err
}

func (s *ScaletService) Create(makeFrom, rplan, name, password, location string, doStart bool,
	keys []int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

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

	res, err := s.client.ExecuteRequest("POST", "scalets", b, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) CreateWithoutPassword(makeFrom, rplan, name, location string, doStart bool,
	keys []int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

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
		Location string  `json:"location,omitempty"`
	}{makeFrom, rplan, doStart, name, keys, location}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("POST", "scalets", b, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Remove(CTID int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("DELETE", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Get(CTID int64) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	res, err := s.client.ExecuteRequest("GET", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), []byte{}, scalet)

	return scalet, res, err
}

func (s *ScaletService) Restart(CTID int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/restart"), []byte{}, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Rebuild(CTID int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/rebuild"), []byte{}, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Stop(CTID int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/stop"), []byte{}, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Start(CTID int64, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/start"), []byte{}, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Upgrade(CTID int64, rplan string, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

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

	res, err := s.client.ExecuteRequest("POST", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10), "/upgrade"), b, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}

func (s *ScaletService) Tasks() (*[]Task, *http.Response, error) {

	tasks := new([]Task)

	res, err := s.client.ExecuteRequest("GET", "tasks", []byte{}, tasks)

	return tasks, res, err
}

func (s *ScaletService) AddSSHKeys(CTID int64, keys []int64) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	body := struct {
		Keys []int64 `json:"keys,omitempty"`
	}{keys}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprint("scalets/", strconv.FormatInt(CTID, 10)), b, scalet)

	return scalet, res, err
}

func (s *ScaletService) Backup(CTID int64, name string) (*Backup, *http.Response, error) {

	backup := new(Backup)

	body := struct {
		Name string `json:"name,omitempty"`
	}{name}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("POST", fmt.Sprintf("scalets/%d/backup", CTID), b, backup)

	return backup, res, err
}

func (s *ScaletService) Restore(CTID int64, makeFrom string, wait bool) (*Scalet, *http.Response, error) {

	scalet := new(Scalet)

	conn := new(websocket.Conn)
	var wsserr error

	if wait {
		conn, wsserr = s.client.WSSConn()
		defer conn.Close()
	}

	body := struct {
		MakeFrom string `json:"make_from,omitempty"`
	}{makeFrom}

	b, _ := json.Marshal(body)

	res, err := s.client.ExecuteRequest("PATCH", fmt.Sprintf("scalets/%d/rebuild", CTID), b, scalet)

	if wait && wsserr == nil && res.Header.Get("VSCALE-TASK-ID") != "" {
		_, err := s.client.WaitTask(conn, res.Header.Get("VSCALE-TASK-ID"))
		return scalet, res, err
	}

	return scalet, res, err
}
