package client

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

type Client struct {
	client  http.Client
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	return &Client{
		client: http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseUrl,
	}
}

func (cl *Client) SendGaugeMetric(name string, value float64) error {
	urlPath := path.Join("/gauge", name, fmt.Sprintf("%f", value))
	request, err := http.NewRequest(http.MethodPost, cl.baseUrl+urlPath, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Set("Content-Type", "text/plain")
	response, err := cl.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()

	_, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (cl *Client) SendCounterMetric(name string, value int64) error {
	urlPath := path.Join("/counter", name, fmt.Sprintf("%d", value))
	request, err := http.NewRequest(http.MethodPost, cl.baseUrl+urlPath, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Set("Content-Type", "text/plain ")
	response, err := cl.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()
	_, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
