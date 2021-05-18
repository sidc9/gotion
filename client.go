package gotion

type Client struct {
	apiKey       string
	baseURL      string
	responseFile string
}

const DefaultURL = "https://api.notion.com/v1"

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = DefaultURL
	}

	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *Client) SaveResponse(filename string) {
	c.responseFile = filename
}
