package mlb

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	DefaultBaseURL = &url.URL{
		Host:   "statsapi.mlb.com",
		Scheme: "https",
		Path:   "/api/v1",
	}
)

type Client struct {
	BaseUrl    *url.URL
	HTTPClient *http.Client
}

func NewDefaultClient() *Client {
	return &Client{
		BaseUrl:    DefaultBaseURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s status code: %d", req.URL.String(), res.StatusCode)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) Get(endpoint string, query map[string]string) ([]byte, error) {
	// Build base http request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.BaseUrl.String(), endpoint), nil)
	if err != nil {
		return nil, err
	}

	// Add query params for request
	q := req.URL.Query()
	q.Add("sportId", "1")

	for key, val := range query {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()

	// Do the request
	b, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return b, nil
}
