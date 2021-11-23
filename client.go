package goinvest

import (
	"encoding/json"
	"log"

	"github.com/foolin/pagser"
)

// NewClient - create new Client
func NewClient() *Client {
	return &Client{}
}

// GetScreener data
func (c *Client) GetScreener(task ScreenerTask) (*ScreenerResponse, error) {
	url := c.buildScreenerRequestURL(task)
	pageBytes, err := c.get(url)
	if err != nil {
		return nil, err
	}

	// new default config
	p := pagser.New()

	// data parser model
	var data screenerPageData

	// parse html data
	err = p.Parse(&data, string(pageBytes))
	// check error
	if err != nil {
		return nil, err
	}

	// print data
	log.Printf("Page data json: \n-------------\n%v\n-------------\n", toJSON(data))

	return nil, nil
}

func (c *Client) buildScreenerRequestURL(task ScreenerTask) string {
	// TODO: use task
	return "https://www.investing.com/stock-screener/?sp=country::56|sector::a|industry::a|equityType::a|exchange::40|eq_pe_ratio::1.94,7|ttmpr2rev_us::0.03,2|price2bk_us::0.01,1|qtotd2eq_us::0,91.63|yield_us::4,14%3Ceq_market_cap;1"
}

func toJSON(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}
