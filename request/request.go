package request

import (
	"encoding/json"
	"net/url"
)

type Interface interface {
	URL() *url.URL
	Path() string
	Body() string
	Query() Query
	Method() string
	RawURL() string
	Expires() int
	Headers() Headers
	json.Unmarshaler
}

type Request struct {
	// UniversalResourceLocator *url.URL

	url     string
	body    string
	method  string
	expires int
	headers Headers
}

func (r *Request) Path() string {
	return r.parsePathQuery().Path
}

func (r *Request) URL() *url.URL {
	u, _ := url.Parse(r.url)
	return u
}

func (r *Request) RawURL() string {
	return r.url
}

func (r *Request) Headers() Headers {
	return r.headers
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) Body() string {
	return r.body
}

func (r *Request) Expires() int {
	return r.expires
}