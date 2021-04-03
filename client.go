package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jordanorelli/blammo"
)

type client struct {
	*blammo.Log
	token string
}

func (c *client) auth(req *http.Request) {
	req.Header.Add("Accept", accept)
	req.Header.Add("Authorization", "token "+options.token)
}

func (c *client) newRequest(method, path string, body io.Reader) *http.Request {
	url := url.URL{
		Scheme: "https",
		Path:   path,
		Host:   "api.github.com",
	}
	req, _ := http.NewRequest(method, url.String(), body)
	c.auth(req)

	return req
}

func (c *client) gists(page int) ([]gist, error) {
	var q url.Values
	if page > 0 {
		q = url.Values{"page": {strconv.Itoa(page)}}
	}
	req := c.newRequest("GET", "/gists", nil)
	req.URL.RawQuery = q.Encode()

	c.Info("GET %s", req.URL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to request gists: %w", err)
	}

	var gists []gist
	if err := json.NewDecoder(res.Body).Decode(&gists); err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}
	return gists, nil
}
