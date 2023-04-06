package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
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
	HttpClientDriver *http.Client
	TargetServer     string
	VictimServer     string
}

func NewDefaultClient() *Client {
	return NewClient("default", &http.Client{}, "http://localhost:8080")
}

func NewClient(name string, clientDriver *http.Client, targetServer string) *Client {
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
		c := NewClient("default", &http.Client{}, "http://localhost:8080")
		return c
	}
	return instance
}

func (c *Client) Get(url string, params string) (interface{}, error) {
	resp, err := c.HttpClientDriver.Get(c.TargetServer + url + params)
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
	resp, err := c.Post("/camp", jc)
	if err != nil {
		return nil, err
	}
	c.VictimServer = resp.(string)
	return resp, nil
}

func (c *Client) GetCampInfo() (interface{}, error) {
	return c.Get("/camp", "")
}

func (c *Client) Ping() (interface{}, error) {
	return c.Get("/ping", "")
}

func (c *Client) ReceiveOrder() (interface{}, error) {
	return c.Get("/order", "?name="+c.Name)
}

func (c *Client) ListenToOrders() {
	var pervOder string
	for {
		order, err := c.ReceiveOrder()
		if err != nil {
			panic(err)
		}
		if pervOder == order {
			continue
		}
		if order == ATTACK {
			log.Printf("ATTACKING! DDOS attack on %s", c.VictimServer)
		}
		if order == STOP {
			log.Printf("DDOS attack on %s stopped", c.TargetServer)
		}
		if order == NOTHING {

		}
		time.Sleep(2 * time.Second)
	}
}

func (c *Client) LeaveCamp() (interface{}, error) {
	return c.Get("/leave", "")
}

func (c *Client) MakeOrder(order string, secretCode string) (interface{}, error) {
	return c.Post("/order", order)
}
