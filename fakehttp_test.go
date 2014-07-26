package fakehttp_test

import (
	"io"
	"net/http"
	"net/url"
	"testing"

	. "github.com/nowk/go-fakehttp"
)

type tHttpClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
}

type FakeApi struct {
	HttpClient tHttpClient
}

type FakeError struct {
	err string
}

func (e *FakeError) Error() string {
	return e.err
}

func TestFakeResponse(t *testing.T) {
	mock := RespondWith(t, 200, "", nil)
	api := FakeApi{mock}
	res, _ := api.HttpClient.Get("http://example.com")

	if res.StatusCode != 200 {
		t.Errorf("Expected `StatusCode` to be 200, got %d", res.StatusCode)
	}
}

func TestDo(t *testing.T) {
	url, _ := url.Parse("http://example.com")
	req := &http.Request{
		Method: "GET",
		URL:    url,
	}

	mock := RespondWith(t, 404, `{"error": "page not found"}`, nil)
	api := FakeApi{mock}
	api.HttpClient.Do(req)

	mock.AssertRequestUrl("http://example.com")
}

func TestError(t *testing.T) {
	mock := ErrorOut(t, &FakeError{"Boom!"})
	api := FakeApi{mock}
	res, err := api.HttpClient.Get("http://example.com")

	if res != nil {
		t.Error("Expected response to be `nil`")
	}

	expectedError := "Boom!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message to be `%s`, but was `%s`", expectedError, err.Error())
	}
}
