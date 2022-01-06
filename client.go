package goinvest

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	screenerAPIURL = "https://www.investing.com/stock-screener/Service/SearchStocks"
)

var (
	countryCodes = map[string]string{
		"USA":    "5",
		"Russia": "56",
	}
	exchangeCodes = []string{
		"2",  // NASDAQ
		"1",  // NYSE
		"50", // NYSE Amex
		// "95", // OTC Markets
	}
)

// NewClient - create new Client
func NewClient() *Client {
	return &Client{}
}

// GetScreener data
func (c *Client) GetScreener(task ScreenerTask) (*ScreenerResponse, error) {
	reqData, err := c.buildRequestData(task)
	if err != nil {
		return nil, err
	}

	response, err := c.POST(screenerAPIURL, reqData)
	if err != nil {
		return nil, err
	}

	result := ScreenerResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.New("failed to decode response as json to struct: " + err.Error())
	}
	return &result, nil
}

func (c *Client) getStockCountryCode(country string) (string, error) {
	code, isCountryFound := countryCodes[country]
	if !isCountryFound {
		return "", errors.New("country `" + country + "` doesn't exists in predefined list")
	}
	return code, nil
}

func (c *Client) getStockSector(sector string) string {
	if sector == "" || sector == "all" {
		return "7,5,21,12,3,16,8,17,13,9,19,1,6,18,15,20,14,23,2,4,10,11,22" // all sectors
	}
	return sector
}

func (c *Client) getStockIndustry(industry string) string {
	if industry == "" || industry == "all" {
		return "81,56,110,59,119,41,120,68,67,88,124,125,51,72,147,136,47,12" +
			",144,8,50,111,2,151,71,9,105,69,45,117,156,46,13,94,102,95,58,100" +
			",101,87,31,106,6,38,112,150,79,107,30,77,131,130,149,160,113,165,28" +
			",158,5,103,163,170,60,18,26,137,135,44,35,53,166,48,141,49,142,143,55" +
			",129,126,139,169,114,153,78,7,86,10,164,132,1,34,154,3,127,146,115,11," +
			"121,162,62,16,108,24,20,54,33,83,29,152,76,133,167,37,90,85,82,104,22,14" +
			",17,109,19,43,140,89,145,96,57,84,118,93,171,27,74,97,4,73,36,42,98,65,70," +
			"40,99,39,92,122,75,66,63,21,159,25,155,64,134,157,128,61,148,32,138,91,116," +
			"123,52,23,15,80,168,161" // all industrys
	}
	return industry
}

func (c *Client) buildRequestData(task ScreenerTask) (url.Values, error) {
	countryCode, err := c.getStockCountryCode(task.Country)
	if err != nil {
		return url.Values{}, err
	}

	result := url.Values{
		"exchange[]": exchangeCodes,
	}

	result.Set("country[]", countryCode)
	result.Set("sector", c.getStockSector(task.Sector))
	result.Set("industry", c.getStockIndustry(task.Industry))
	result.Set("equityType", "ORD") // only ordinary shares

	setTaskBarValue(&result, "eq_pe_ratio", task.PriceToEarnings)   // P/E
	setTaskBarValue(&result, "ttmpr2rev_us", task.PriceToSales)     // P/S
	setTaskBarValue(&result, "price2bk_us", task.PriceToBook)       // P/B
	setTaskBarValue(&result, "qtotd2eq_us", task.TotalDebtToEquity) // Total Debt / Equity
	setTaskBarValue(&result, "last", task.LastRate)                 // last price
	setTaskBarValue(&result, "eq_dividend", task.Dividend)          // dividend
	setTaskBarValue(&result, "yield_us", task.DividendYield)        // dividend yield

	result.Set("pn", "1")                // ??
	result.Set("order[col]", "yield_us") // order by
	result.Set("order[dir]", "a")        // order direction

	return result, nil
}

func setTaskBarValue(result *url.Values, field string, value *ValueBar) {
	if value == nil {
		return
	}
	result.Set(field+"[min]", floatToString(value.Min))
	result.Set(field+"[max]", floatToString(value.Max))
}
