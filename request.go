package goinvest

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// get - send http GET request
func (c *Client) get(url string) ([]byte, error) {
	// declare HTTP Method and Url
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to send GET request: " + err.Error())
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
