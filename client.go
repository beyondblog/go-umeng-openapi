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
	Apps *Apps
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	c.Apps = &Apps{c}
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

	fmt.Println(url, c.Authorization, form.Encode())

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

	fmt.Println(url, c.Authorization, form.Encode())

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
