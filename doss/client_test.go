package doss

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	//first create client
	NewClient("Mohammad", http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	})
	//must throw exception when new client created
	t.Run("already", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("must panic because of creation 2 instance of clients")
			}
		}()
		NewClient("Hello", http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		})
	})
}
