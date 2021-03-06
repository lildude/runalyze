package runalyze

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

var (
	// BaseURLV1 is Runalyze's v1 API endpoint
	BaseURLV1 = "https://runalyze.com/api/v1/"
	userAgent = "go-runalyze"
)

// Client holds configuration items for the Runalyze client and provides methods that interact with the Runalyze API.
type Client struct {
	baseURL   *url.URL
	apiToken  string
	UserAgent string
	client    *http.Client
}

// NewClient returns a new Runalyze API client.
func NewClient(token string) *Client {
	cc := http.DefaultClient
	baseURL, _ := url.Parse(BaseURLV1)

	c := &Client{baseURL: baseURL, UserAgent: userAgent, apiToken: token, client: cc}
	return c
}

// NewRequest creates an HTTP Request. The client baseURL is checked to confirm that it has a trailing
// slash. A relative URL should be provided without the leading slash. If a non-nil body is provided
// it will be JSON encoded and included in the request.
// Inspiration: https://github.com/google/go-github/blob/master/github/github.go
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("client baseURL does not have a trailing slash: %q", c.baseURL)
	}

	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("token", c.apiToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends a request and returns the response. An error is returned if the request cannot
// be sent or if the API returns an error. If a response is received, the response body
// is decoded and stored in the value pointed to by v.
// Inspiration: https://github.com/google/go-github/blob/master/github/github.go
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(err, ctx.Err().Error())
		default:
			return nil, err
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read body")
	}
	resp.Body.Close()

	// Anything other than a HTTP 2xx response code is treated as an error.
	if c := resp.StatusCode; c >= 300 {

		// Handle auth errors
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			err := AuthError(http.StatusText(resp.StatusCode))
			return resp, err
		}

		// Try parsing the response using the standard error schema. If this fails we wrap the parsing
		// error and return. Otherwise return the errors included in the API response payload.
		var e Error
		err := json.Unmarshal(data, &e)
		if err != nil {
			err = errors.Wrap(err, http.StatusText(resp.StatusCode))
			return resp, errors.Wrap(err, "unable to parse API error response")
		}

		return resp, errors.Wrap(e, http.StatusText(resp.StatusCode))
	}

	if v != nil && len(data) != 0 {
		err = json.Unmarshal(data, v)

		switch err {
		case nil:
		case io.EOF:
			err = nil
		default:
			err = errors.Wrap(err, "unable to parse API response")
		}
	}

	return resp, err
}
