package anthropic

// Client adapts Anthropic Messages API.
type Client struct {
	BaseURL string
	APIKey  string
}

func New(baseURL, apiKey string) *Client {
	return &Client{BaseURL: baseURL, APIKey: apiKey}
}