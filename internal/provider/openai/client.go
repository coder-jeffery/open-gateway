package openai

// Client adapts OpenAI-compatible HTTP APIs.
type Client struct {
	BaseURL string
	APIKey  string
}

func New(baseURL, apiKey string) *Client {
	return &Client{BaseURL: baseURL, APIKey: apiKey}
}