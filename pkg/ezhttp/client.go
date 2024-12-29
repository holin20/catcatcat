package ezhttp

import (
	"io"
	"net/http"
)

type httpClient struct {
	headers map[string]string
	client  *http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		headers: make(map[string]string),
		client:  &http.Client{},
	}
}

func NewHttpClientWithCustomClient(client *http.Client) *httpClient {
	return &httpClient{
		headers: make(map[string]string),
		client:  client,
	}
}

func (c *httpClient) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *httpClient) WithDefaultUserAgent() *httpClient {
	c.SetHeader(headerUserAgent, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	return c
}

func (c *httpClient) Get(url string) (string, error) {
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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
