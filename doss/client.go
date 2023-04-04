package doss

import (
	"DDOS_ARMY/antena"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Name             string      `json:"NAME"`
	HttpClientDriver http.Client `json:"httpClient"`
}

var instance *Client

func isInstantiated() bool {
	return instance != nil
}

func NewDefaultClient() Client {
	return NewClient("Ali", http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	})
}

func NewClient(name string, clientDriver http.Client) Client {
	if isInstantiated() {
		panic("client already exists")
	}
	instance = &Client{
		Name:             name,
		HttpClientDriver: clientDriver,
	}
	return *instance
}

func GetClient() Client {
	if !isInstantiated() {
		panic("client not instantiated")
	}
	return *instance
}

func (cl Client) Ping(target string, port int) antena.Response {
	if !isInstantiated() {
		panic("client not instantiated")
	}

	req, err := http.NewRequest(http.MethodGet, target+":"+strconv.Itoa(port)+"/ping", nil)
	if err != nil {
		panic(err)
	}
	do, err := cl.HttpClientDriver.Do(req)
	if err != nil {
		return antena.Response{}
	}
	var pongResponse antena.Response
	err = json.NewDecoder(do.Body).Decode(&pongResponse)
	if err != nil {
		return antena.Response{}
	}
	defer do.Body.Close()
	return pongResponse
}
