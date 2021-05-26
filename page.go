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

type PageContent struct {
	Object     string   `json:"object"`
	Results    []*Block `json:"results"`
	HasMore    bool     `json:"has_more"`
	NextCursor string   `json:"next_cursor"`
}

// TODO get block children
type Block struct {
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Object         string `json:"object"`
	HasChildren    bool   `json:"has_children"`
	Type           string `json:"type"`

	Heading1  BlockText `json:"heading_1"`
	Heading2  BlockText `json:"heading_2"`
	Heading3  BlockText `json:"heading_3"`
	Paragraph BlockText `json:"paragraph"`
	// TODO: more types of content
}

type BlockText struct {
	Text []*RichText `json:"text"`
}

func (b *Block) Content() (interface{}, error) {
	switch b.Type {
	case "heading_1":
		return b.Heading1, nil
	case "heading_2":
		return b.Heading2, nil
	case "heading_3":
		return b.Heading3, nil
	case "paragraph":
		return b.Paragraph, nil
	}
	return nil, fmt.Errorf("block type is not set")
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
