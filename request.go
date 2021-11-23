package goinvest

import (
	"errors"
	"fmt"
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

	// check response code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
