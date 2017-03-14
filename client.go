package umeng

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	*Config
	Apps   *Apps
	Events *Events
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	c.Apps = &Apps{c}
	c.Events = &Events{c}
	return c
}

// get rpc style endpoint.
func (c *Client) get(path string, in interface{}) (io.ReadCloser, error) {

	form, _ := query.Values(in)

	url := "http://api.umeng.com" + path + "?" + form.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+c.Authorization)

	return c.do(req)
}

// post rpc style endpoint.
func (c *Client) post(path string, in interface{}) (io.ReadCloser, error) {
	url := "http://api.umeng.com" + path

	form, _ := query.Values(in)

	req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+c.Authorization)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.do(req)
}

// perform the request.
func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		defer res.Body.Close()
		if b, err := ioutil.ReadAll(res.Body); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("response %s: %s", res.Status, b)
		}
	}

	return res.Body, err
}
