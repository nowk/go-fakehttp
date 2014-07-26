package fakehttp

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// HttpClient
type HttpClient struct {
	StatusCode int
	Body       io.ReadCloser
	Header     http.Header

	t      *testing.T
	reqUrl string
	err    error
}

// New returns new fake http client to be configured
func New(t *testing.T) *HttpClient {
	client := new(HttpClient)
	client.t = t
	return client
}

// RespondWith returns a fake http client that will respond with given status,
// body and headers when request is made on the mocked http.Client
func RespondWith(t *testing.T, status int, body string, headers http.Header) *HttpClient {
	client := New(t)
	client.StatusCode = status
	client.Header = headers
	client.Body = ReadCloser{strings.NewReader(body)}
	return client
}

// ErrorOut returns the error for an http.DefaultClient method (..., error)
// This is not for not-2xx response
func ErrorOut(t *testing.T, err error) *HttpClient {
	client := New(t)
	client.err = err
	return client
}

// GetReqUrl returns the requested URL
func (f HttpClient) GetReqUrl() string {
	return f.reqUrl
}

// methods mapping to http.Client
// saving particulars to HttpClient for later use, eg. AssertRequestUrl

func (f *HttpClient) Do(req *http.Request) (*http.Response, error) {
	return f.fakeResponse(req.URL.String())
}

func (f *HttpClient) Get(urlStr string) (*http.Response, error) {
	return f.fakeResponse(urlStr)
}

func (f *HttpClient) Head(urlStr string) (*http.Response, error) {
	return f.fakeResponse(urlStr)
}

func (f *HttpClient) Post(urlStr string, bodyType string, body io.Reader) (*http.Response, error) {
	return f.fakeResponse(urlStr)
}

func (f *HttpClient) PostForm(urlStr string, data url.Values) (*http.Response, error) {
	return f.fakeResponse(urlStr)
}

// fakeResponse returns a mocked response with defined properties
func (f *HttpClient) fakeResponse(urlStr string) (*http.Response, error) {
	f.reqUrl = urlStr

	if f.err != nil {
		return nil, f.err
	}

	res := &http.Response{
		StatusCode: f.StatusCode,
		Header:     f.Header,
		Body:       f.Body,
	}
	return res, nil
}

// AssertRequestUrl assertion helper to assert the last requested URL
func (f HttpClient) AssertRequestUrl(urlStr string) {
	if f.reqUrl != urlStr {
		f.t.Errorf("Expected URL to be `%s`, but was `%s`", f.reqUrl, urlStr)
	}
}

// ReadCloser implements io.ReadCloser interface
type ReadCloser struct {
	io.Reader
}

func (r ReadCloser) Close() error {
	return nil
}
