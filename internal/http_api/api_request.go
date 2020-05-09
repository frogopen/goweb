package http_api

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	c *http.Client
}

//data = strings.NewReader("username=1&password=2")
func (c *Client) POST(url string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/x-www-form-urlencoded")

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
