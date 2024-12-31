package ezgo

import (
	"fmt"
	"strings"
)

type Url struct {
	// Required
	scheme string
	domain string

	// Optional
	port uint16
	path string

	queryString      string
	queryParams      map[string]string
	queryStringPatch string
}

func NewUrl(
	scheme string,
	domain string,
) *Url {
	return &Url{
		scheme:      scheme,
		domain:      domain,
		queryParams: make(map[string]string),
	}
}

func NewHttpsUrl(domain string) *Url {
	return NewUrl("https", domain)
}

func NewHttpUrl(domain string) *Url {
	return NewUrl("http", domain)
}

func (u *Url) WithPort(port uint16) *Url {
	u.port = port
	return u
}

func (u *Url) WithPath(path string) *Url {
	u.path = path
	return u
}

func (u *Url) WithQueryString(queryString string) *Url {
	u.queryString = queryString
	return u
}

func (u *Url) WithQueryStringPatch(queryStringPatch string) *Url {
	u.queryStringPatch = queryStringPatch
	return u
}

func (u *Url) String() string {
	return fmt.Sprintf(
		"%s://%s%s%s%s",
		u.scheme,
		u.domain,
		If(u.port > 0, fmt.Sprintf(":%d", u.port), ""),
		If(u.path != "", "/"+u.path, ""),
		If(u.GetQueryString() != "", "?"+u.GetQueryString(), ""),
	)
}

func (u *Url) WithQueryParam(key, value string) *Url {
	u.queryParams[key] = value
	return u
}

func (u *Url) GetQueryString() string {
	resolved := u.queryString
	if resolved == "" {
		resolved = strings.Join(
			FlattenMap(u.queryParams, func(k string, v string) string {
				return k + "=" + v
			}),
			"&",
		)
	}
	return strings.Join(FilterEmpty([]string{resolved, u.queryStringPatch}), "&")
}
