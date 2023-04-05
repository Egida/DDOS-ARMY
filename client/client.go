package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

const (
	ATTACK  = "ATTACK"
	STOP    = "STOP"
	NOTHING = "NOTHING"
)

var instance *Client
var once sync.Once

type JsonClient struct {
	Name string `json:"name"`
}

type Client struct {
	Name             string
	HttpClientDriver http.Client
	TargetServer     string
}

func NewDefaultClient() *Client {
	return NewClient("default", http.Client{}, "http://localhost:8080")
}

func NewClient(name string, clientDriver http.Client, targetServer string) *Client {
	once.Do(func() {
		instance = &Client{
			Name:             name,
			HttpClientDriver: clientDriver,
			TargetServer:     targetServer,
		}
	})
	return instance
}

func GetClient() *Client {
	if instance == nil {
		c := NewClient("default", http.Client{}, "http://localhost:8080")
		return c
	}
	return instance
}

func (c *Client) Get(url string) (interface{}, error) {
	resp, err := c.HttpClientDriver.Get(c.TargetServer + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//if content type is not json, return string
	if resp.Header.Get("Content-Type") != "application/json" {
		jr, err := io.ReadAll(resp.Body)
		return string(jr), err
	}
	var jr interface{}
	err = json.NewDecoder(resp.Body).Decode(&jr)
	return jr, err
}

func (c *Client) Post(url string, data interface{}) (interface{}, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.HttpClientDriver.Post(c.TargetServer+url, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//if the response is not json, return it as string
	if resp.Header.Get("Content-Type") != "application/json" {
		jr, err := io.ReadAll(resp.Body)
		return string(jr), err
	}
	var jr interface{}
	err = json.NewDecoder(resp.Body).Decode(&jr)
	return jr, err
}

func (c *Client) JoinCamp() (interface{}, error) {
	jc := JsonClient{Name: c.Name}
	return c.Post("/camp", jc)
}

func (c *Client) GetCampInfo() (interface{}, error) {
	return c.Get("/camp")
}

func (c *Client) Ping() (interface{}, error) {
	return c.Get("/ping")
}

func (c *Client) ReceiveOrder() (interface{}, error) {
	return c.Get("/order")
}

func (c *Client) ListenToOrders() {
	for {
		order, err := c.ReceiveOrder()
		if err != nil {
			panic(err)
		}
		if order == ATTACK {
			// start DDOS attack
		}
		if order == STOP {
			// stop DDOS attack
		}
		if order == NOTHING {
			// do nothing
		}
	}
}

func (c *Client) LeaveCamp() (interface{}, error) {
	return c.Get("/leave")
}
