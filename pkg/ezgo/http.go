package ezgo

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

type HttpClient struct {
	headers      map[string]string
	client       *http.Client
	useCookieJar bool
}

func NewHttpClient(useCookieJar bool) *HttpClient {
	return &HttpClient{
		headers: make(map[string]string),
		client: &http.Client{
			Jar: Arg1(cookiejar.New(nil)),
		},
		useCookieJar: useCookieJar,
	}
}

func NewHttpClientWithCustomClient(client *http.Client, useCookieJar bool) *HttpClient {
	if useCookieJar && client.Jar == nil {
		client.Jar = Arg1(cookiejar.New(nil))
	}
	return &HttpClient{
		headers:      make(map[string]string),
		client:       client,
		useCookieJar: useCookieJar,
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

func (c *HttpClient) Get(url string) (string, error) {
	req, err := http.NewRequest(methodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Set custom headers
	for key, value := range c.headers {
		if key == headerCookie &&
			c.useCookieJar &&
			c.client.Jar != nil &&
			len(c.client.Jar.Cookies(req.URL)) > 0 {
			// we will handle cookies using jar instead
			continue
		}
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp != nil && c.useCookieJar && c.client.Jar != nil {
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
