package goinvest

// Client - investing.com data client
type Client struct{}

// ScreenerTask - the task to retrieve the screener data
type ScreenerTask struct {
	Country  string `json:"country"`  // example: "USA"
	Sector   string `json:"sector"`   // example: "all"
	Industry string `json:"industry"` // example: "all"
}

// ScreenerResponse data
type ScreenerResponse struct {
	Data []StockData `json:"hits"`
}

// StockData - share-specific data
type StockData struct {
	Name       string        `json:"name_trans"`
	PairID     int64         `json:"pair_ID"`
	Symbol     string        `json:"stock_symbol"`
	ShareType  string        `json:"security_type"` // example: "ORD" for ordinary shares
	IndustryID int64         `json:"industry_id"`
	ShareData  StockPairData `json:"viewData"`

	// sector & industry
	SectorID     int64  `json:"sector_id"`
	SectorName   string `json:"sector_trans"`
	IndustryInfo string `json:"industry_trans"`

	// exchange
	ExchangeID   int64  `json:"exchange_ID"`
	ExchangeName string `json:"exchange_trans"`

	// fundamental
	PriceToEarnings   float64 `json:"eq_pe_ratio"`   // P/E
	MarketCap         float64 `json:"eq_market_cap"` // market cap
	Dividend          float64 `json:"eq_dividend"`
	DividendYield     string  `json:"yield_us"`
	Revenue           float64 `json:"eq_revenue"`
	PriceToSales      float64 `json:"ttmpr2rev_us"` // P/S
	PriceToBook       float64 `json:"price2bk_us"`  // P/B
	TotalDebtToEquity float64 `json:"qtotd2eq_us"`

	// formated data
	PriceToEarningsFormated   string `json:"eq_pe_ratio_frmt"`
	MarketCapFormated         string `json:"eq_market_cap_frmt"`
	RevenueFormated           string `json:"eq_revenue_frmt"` // example: "606.01M"
	TotalDebtToEquityFormated string `json:"qtotd2eq_us_eu"`

	//indicators
	RSI       float64 `json:"RSI"`
	STOCH     float64 `json:"STOCH"`
	CCI       float64 `json:"CCI"`
	MACD      float64 `json:"MACD"`
	ADX       float64 `json:"ADX"`
	WilliamsR float64 `json:"WilliamsR"`
	STOCHRSI  float64 `json:"STOCHRSI"`
	ATR       float64 `json:"ATR"`
	HL        float64 `json:"HL"`
	UO        float64 `json:"UO"`
	ROC       float64 `json:"ROC"`
	BullBear  float64 `json:"BullBear"`

	// tech-indicators stock forecast
	TechSumWeek  string `json:"tech_sum_week"`
	TechSumMonth string `json:"tech_sum_month"`

	// rate
	LastRate        float64 `json:"last"`
	ChangeYearToDay float64 `json:"ytd"`
	ChangeWeek      float64 `json:"week"`
	ChangeMonth     float64 `json:"month"`
	ChangeYear      float64 `json:"year"`
	Change3Years    float64 `json:"3year"`
}

// StockPairData - stock pair data
type StockPairData struct {
	Flag   string `json:"flag"`
	Symbol string `json:"symbol"`
	Link   string `json:"link"`
	Name   string `json:"name"`
}

type dataMap map[string]interface{}
