package main

type Client struct {
	apiKey  string
	baseURL string
}

const defaultURL = "https://api.notion.com/v1"

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = defaultURL
	}

	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}
