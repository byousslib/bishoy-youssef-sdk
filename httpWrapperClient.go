package lotrApiSdk

import (
	"io"
	"net/http"
	"net/url"
)

// this is a helper, unexported struct to wrap the http client and make calls to the API
// easier. This also takes into consideration the base url, so the main package doesn't
// have to create every request using the URL
type httpClient struct {
	c        http.Client
	baseUrl  string
	apiToken string
}

// wrapper for the Get function
func (c *httpClient) Get(requestedPath string) (resp *http.Response, err error) {
	combinedUrl, err := url.JoinPath(c.baseUrl, requestedPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", combinedUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Wrapper for Post functions
func (c *httpClient) Post(requestedPath, contentType string, body io.Reader) (resp *http.Response, err error) {
	combinedUrl, err := url.JoinPath(c.baseUrl, requestedPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", combinedUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// A generic wrapper that is used for any request.
// including the above Get and Post requests
func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.apiToken)
	return c.c.Do(req)
}
