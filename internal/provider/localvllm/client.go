package localvllm

// Client adapts local vLLM / OpenAI-compatible endpoints.
type Client struct {
	BaseURL string
}

func New(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}