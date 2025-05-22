package net

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	TIMEOUT = 10 * time.Second
)

type H map[string]interface{}

type Response struct {
	Data  interface{}
	Error error
}

type cli struct {
	client *http.Client
	body   io.Reader
	header H
}

func NewRequest() *cli {
	return &cli{
		header: nil, body: nil,
		client: &http.Client{
			Timeout: TIMEOUT,
		},
	}
}

func (c *cli) Header(val H) *cli {
	c.header = val
	return c
}

func (c *cli) Timeout(val time.Duration) *cli {
	c.client.Timeout = val
	return c
}

func (c *cli) Proxy(proxyUrl string) *cli {
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.client.Transport = tr
	return c
}

func (c *cli) Body(val H) *cli {
	b, _ := json.Marshal(val)
	reader := bytes.NewReader(b)
	c.body = reader
	return c
}

func (c *cli) handler(url, method string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, c.body)
	if err != nil {
		return
	}
	if c.header != nil {
		for k, v := range c.header {
			req.Header.Set(k, fmt.Sprintf("%v", v))
		}
	}
	response, err = c.client.Do(req)
	if err != nil {
		return
	}
	return
}

func (c *cli) Get(url string) (r Response) {
	response, err := c.handler(url, GET)
	if err != nil {
		r.Error = err
		return
	}
	defer response.Body.Close()

	r.Data, r.Error = io.ReadAll(response.Body)
	return
}

func (c *cli) GetStatusCode(url string) (r Response) {
	response, err := c.handler(url, GET)
	if err != nil {
		r.Error = err
		return
	}
	defer response.Body.Close()

	r.Data = response.StatusCode
	return
}

func (c *cli) HeadStatusCode(url string) (r Response) {
	response, err := c.handler(url, HEAD)
	if err != nil {
		r.Error = err
		return
	}
	defer response.Body.Close()

	r.Data = response.StatusCode
	return
}

func (c *cli) Post(url string) (r Response) {
	response, err := c.handler(url, POST)
	if err != nil {
		r.Error = err
		return
	}
	defer response.Body.Close()

	r.Data, r.Error = io.ReadAll(response.Body)
	return
}

func (c *cli) PostJSON(url string) (r Response) {
	c.header = H{"Content-Type": "application/json"}
	r = c.Post(url)
	return
}

func (c *cli) PostForm(url string) (r Response) {
	c.header = H{"Content-Type": "application/x-www-form-urlencoded"}
	r = c.Post(url)
	return
}
