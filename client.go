package goinvest

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// NewClient - create new Client
func NewClient() *Client {
	return &Client{}
}

// GetScreener data
func (c *Client) GetScreener(task ScreenerTask) (*ScreenerResponse, error) {
	url := c.buildScreenerRequestURL(task)

	// Load the HTML document
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find("table#resultsTable > tbody > tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td.symbol").Text()
		fmt.Println(title)
	})

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
