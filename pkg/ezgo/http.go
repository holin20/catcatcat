package ezgo

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

type HttpClient struct {
	headers map[string]string
	client  *http.Client
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

func (c *HttpClient) WithDefaultUserAgent() *HttpClient {
	c.SetHeader(headerUserAgent, defaultUserAgent)
	return c
}

func (c *HttpClient) Get(url string, respectToRespCookie bool) (string, error) {
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
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp != nil && respectToRespCookie {
		if c.client.Jar == nil {
			c.client.Jar, _ = cookiejar.New(nil)
		}

		respCookies := resp.Cookies()
		if len(respCookies) > 0 {
			c.client.Jar.SetCookies(req.URL, respCookies)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
