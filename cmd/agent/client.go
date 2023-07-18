package main

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

const BaseURL = "http://localhost:8080/update/"

type Client struct {
	client http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (cl *Client) SendGaugeMetric(name string, value float64) {
	urlPath := path.Join("gauge", name, fmt.Sprintf("%f", value))
	request, err := http.NewRequest(http.MethodPost, BaseURL+urlPath, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("Content-Type", "text/plain")
	response, err := cl.client.Do(request)
	if err != nil {
		fmt.Println(err, "ddd")
		return
	}
	defer response.Body.Close()

	_, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (cl *Client) SendCounterMetric(name string, value int64) {
	urlPath := path.Join("counter", name, fmt.Sprintf("%d", value))
	request, err := http.NewRequest(http.MethodPost, BaseURL+urlPath, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("Content-Type", "text/plain ")
	response, err := cl.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	_, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
