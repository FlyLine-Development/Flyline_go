package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	httpClient *http.Client
}

type ClientOptions struct {
	HTTPClient *http.Client
}

func NewClient(options ClientOptions) (client *Client, err error) {

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

	return &Client{
		httpClient: options.HTTPClient,
	}, nil
}

func (c *Client) Call(endpoint string, body []byte, v interface{}, requestmethod string) error {
	if requestmethod == "GET" {
		req, err := c.newGetRequest(endpoint, bytes.NewReader(body), v)
		if err != nil {
			return err
		}

		return c.do(req, v)
	} else {
		req, err := c.newPostRequest(endpoint, bytes.NewReader(body), v)
		if err != nil {
			return err
		}

		return c.do(req, v)
	}
}

// newRequest is used by Call to generate a http.Request with appropriate headers.
func (c *Client) newPostRequest(endpoint string, body io.Reader, v interface{}) (*http.Request, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) newGetRequest(endpoint string, body io.Reader, v interface{}) (*http.Request, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequest("GET", endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode == 200 {
		return json.NewDecoder(res.Body).Decode(v)
	}

	var flylineErr Error
	if err = json.NewDecoder(res.Body).Decode(&flylineErr); err != nil {
		return err
	}
	flylineErr.StatusCode = res.StatusCode
	return flylineErr
}
