package geo_search

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	baseURL = "http://api.geonames.org/searchJSON?"
)

type Client struct {
	timeout  time.Duration
	maxRows  int
	language string
	username string
}

// Create new client. Username for authentication. Language in ISO-639 format, 2-letter language code (en,de,fr,ud,es,ru)
// Rows sets the maximum number of returned places. Timeout is a timeout...
func New(username, lang string, rows int, timeout time.Duration) *Client {
	return &Client{
		timeout:  timeout,
		maxRows:  rows,
		language: lang,
		username: username,
	}
}

// Search place by some data like place name, country name, continent, admin codes,...
func (c *Client) Search(query string) ([]GeoData, error) {
	escaped := url.QueryEscape(query)
	cl := makeClient(c.timeout)
	resp, err := cl.Get(baseURL + fmt.Sprintf("q=%s&maxRows=%d&lang=%s&username=%s", escaped, c.maxRows, c.language, c.username))
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	data, err := extractData(resp)
	if err != nil {
		return nil, errors.Wrap(err, "data parsing error")
	}
	return data.Parse(), nil
}

// Get place only by name
func (c *Client) ByName(name string) ([]GeoData, error) {
	escaped := url.QueryEscape(name)
	cl := makeClient(c.timeout)
	resp, err := cl.Get(baseURL + fmt.Sprintf("name=%s&maxRows=%d&lang=%s&username=%s", escaped, c.maxRows, c.language, c.username))
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	data, err := extractData(resp)
	if err != nil {
		return nil, errors.Wrap(err, "data parsing error")
	}
	return data.Parse(), nil
}

// Get exact places
func (c *Client) ByNameEquals(name string) ([]GeoData, error) {
	escaped := url.QueryEscape(name)
	cl := makeClient(c.timeout)
	resp, err := cl.Get(baseURL + fmt.Sprintf("name_equals=%s&maxRows=%d&lang=%s&username=%s", escaped, c.maxRows, c.language, c.username))
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	data, err := extractData(resp)
	if err != nil {
		return nil, errors.Wrap(err, "data parsing error")
	}
	return data.Parse(), nil
}

func makeClient(timeout time.Duration) http.Client {
	return http.Client{
		Timeout: timeout,
	}
}
