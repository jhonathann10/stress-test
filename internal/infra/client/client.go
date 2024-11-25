package client

import (
	"net/http"
)

type ClientInterface interface {
	Get() (int, error)
}

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func (c *Client) Get() (int, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
