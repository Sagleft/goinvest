package goinvest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GET - send http GET request
func (c *Client) GET(url string) ([]byte, error) {
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

// POST - send http POST request
func (c *Client) POST(queryURL string, data url.Values) ([]byte, error) {
	// declare http client
	httpClient := &http.Client{}

	// declare HTTP Method and Url
	req, err := http.NewRequest("POST", queryURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.New("failed to send POST request: " + err.Error())
	}

	// req.Header.Set("Cookie", "name=value")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "www.investing.com")
	req.Header.Set("Origin", "https://www.investing.com")
	req.Header.Set("Referer", "https://www.investing.com/stock-screener/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// send request
	resp, err := httpClient.Do(req)

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
