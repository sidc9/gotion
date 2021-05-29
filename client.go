package gotion

import "net/http"

type Client struct {
	apiKey       string
	baseURL      string
	responseFile string
	httpClient   *http.Client
}

const DefaultURL = "https://api.notion.com/v1"

var client *Client

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = DefaultURL
	}

	client = &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
	return client
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

func (c *Client) WithHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}
