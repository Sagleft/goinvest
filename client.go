package goinvest

// NewClient - create new Client
func NewClient() *Client {
	return &Client{}
}

// GetScreener data
func (c *Client) GetScreener(task ScreenerTask) (*ScreenerResponse, error) {
	// TODO
	return nil, nil
}
