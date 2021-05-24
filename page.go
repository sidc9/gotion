package gotion

import (
	"fmt"
	"net/http"
)

type Page struct {
	CreatedTime    string                 `json:"created_time"`
	ID             string                 `json:"id"`
	LastEditedTime string                 `json:"last_edited_time"`
	Object         string                 `json:"object"`
	Properties     PageProperties         `json:"properties"`
	Archived       bool                   `json:"archived"`
	Parent         map[string]interface{} `json:"parent"`
}

func (c *Client) GetPage(id string) (*Page, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var page Page
	err := c.doRequest(http.MethodGet, fmt.Sprintf("pages/%s", id), nil, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

// TODO children, example/test
type requestBody struct {
	Properties map[string]*PageProperty `json:"properties"`
	Children   interface{}              `json:"children"`
}

func (c *Client) CreatePage(parent string, properties map[string]*PageProperty, children interface{}) error {
	if parent == "" {
		return fmt.Errorf("parentID is required")
	}

	body := requestBody{
		Properties: properties,
		Children:   children,
	}

	return c.doRequest(http.MethodPost, "pages", body, nil)
}
