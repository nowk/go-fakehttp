# go-fakehttp

[![Build Status](https://travis-ci.org/nowk/go-fakehttp.svg?branch=master)](https://travis-ci.org/nowk/go-fakehttp)

Fake `http.DefaultClent`, when you can't control the endpoint URLs for testing.

## Example

    type FakeApi struct {
      HttpClient YourHttpClientInterface
    }

    func TestPageNotFound(t *testing.T) {
      mock := fakehttp.RespondWith(t, 404, `{"error": "page not found"}`, nil)
      api := FakeApi{mock}

      res, err := api.HttpClient.Get("http://example.com")
      if err != nil {
        t.Error("Received an error, when we should not have")
      }

      if res.StatusCode != 404 {
        t.Errorf("Expected `StatusCode` to be 404, got %d", res.StatusCode)
      }
    }

### License

MIT
