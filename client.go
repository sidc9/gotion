package gotion

type Client struct {
	apiKey       string
	baseURL      string
	responseFile string
}

const DefaultURL = "https://api.notion.com/v1"

var client *Client

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

func Init(apiKey, baseURL string) {
	client = NewClient(apiKey, baseURL)
}

func SaveResponse(filename string) {
	client.SaveResponse(filename)
}

func GetClient() *Client {
	return client
}
