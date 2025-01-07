package ezgo

import (
	"io"
	"net/http"
	"strings"
)

type HttpClient struct {
	headers map[string]string
	client  *http.Client

	respSetCookie string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		headers: make(map[string]string),
		client:  &http.Client{},
	}
}

func NewHttpClientWithCustomClient(client *http.Client) *HttpClient {
	return &HttpClient{
		headers: make(map[string]string),
		client:  client,
	}
}

func (c *HttpClient) SetHeader(key, value string) *HttpClient {
	c.headers[key] = value
	return c
}

func (c *HttpClient) SetCookieString(value string) *HttpClient {
	c.headers[headerCookie] = value
	return c
}

func (c *HttpClient) SetCookieStringIfNeeded(value string) *HttpClient {
	if c.respSetCookie != "" {
		c.headers[headerCookie] = c.respSetCookie
		return c
	}
	c.headers[headerCookie] = value
	return c
}

func (c *HttpClient) WithDefaultUserAgent() *HttpClient {
	c.SetHeader(headerUserAgent, defaultUserAgent)
	return c
}

func (c *HttpClient) Get(url string, setRespCookie bool) (string, error) {
	req, err := http.NewRequest(methodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Set custom headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp != nil && setRespCookie && len(resp.Header.Values(headerSetCookie)) > 0 {
		c.respSetCookie = strings.Join(resp.Header.Values(headerSetCookie), ";")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
