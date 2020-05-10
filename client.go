package edsm

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseURL = "https://www.edsm.net/"

// Client is a base client
type Client struct {
	UserAgent *string

	httpClient *http.Client
}

// New creates a new EDSM Client
func New(client *http.Client) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) newGetRequest(path string) (*http.Request, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	rel := &url.URL{Path: path}
	u := base.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent != nil {
		req.Header.Set("User-Agent", *c.UserAgent)
	} else {
		req.Header.Set("User-Agent", "go-edsm")
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
