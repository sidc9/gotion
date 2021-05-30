package gotion

import (
	"fmt"
	"net/http"
)

type Page struct {
	ID             string                 `json:"id"`
	CreatedTime    string                 `json:"created_time"`
	LastEditedTime string                 `json:"last_edited_time"`
	Object         string                 `json:"object"`
	Properties     PageProperties         `json:"properties"`
	Archived       bool                   `json:"archived"`
	Parent         map[string]interface{} `json:"parent"`

	c *Client
}

func (p *Page) Title() string {
	for _, v := range p.Properties {
		if v.Type == "title" {
			return v.Title[0].PlainText
		}
	}
	return ""
}

func (p *Page) ParentID() string {
	typ := p.Parent["type"].(string)
	if typ == "workspace" {
		return typ
	}

	return p.Parent[typ].(string)
}

func (p *Page) Content() (*PageContent, error) {
	return p.c.GetPageContent(p.ID)
}

type PageContent struct {
	Object     string   `json:"object"`
	Results    []*Block `json:"results"`
	HasMore    bool     `json:"has_more"`
	NextCursor string   `json:"next_cursor"`
}

func (c *Client) GetPage(id string) (*Page, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	page := &Page{c: c}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("pages/%s", id), nil, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (c *Client) GetPageContent(id string) (*PageContent, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var pc PageContent
	err := c.doRequest(http.MethodGet, fmt.Sprintf("blocks/%s/children", id), nil, &pc)
	if err != nil {
		return nil, err
	}

	return &pc, nil
}

func (c *Client) UpdatePageProperty(pageID, propertyName string, propertyValue PageProperty) (*Page, error) {
	if pageID == "" {
		return nil, fmt.Errorf("page id is required")
	}

	update := map[string]interface{}{
		"properties": map[string]interface{}{
			propertyName: propertyValue,
		},
	}

	page := &Page{c: c}
	err := c.doRequest(http.MethodPatch, fmt.Sprintf("pages/%s", pageID), update, page)
	if err != nil {
		return nil, err
	}

	return page, nil
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
