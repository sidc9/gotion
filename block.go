package gotion

import (
	"fmt"
	"net/http"
)

// TODO get block children
type Block struct {
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Object         string `json:"object"`
	HasChildren    bool   `json:"has_children"`
	Type           string `json:"type"`

	Heading1   BlockText     `json:"heading_1"`
	Heading2   BlockText     `json:"heading_2"`
	Heading3   BlockText     `json:"heading_3"`
	Paragraph  BlockText     `json:"paragraph"`
	BulletList BlockText     `json:"bulleted_list_item"`
	NumberList BlockText     `json:"numbered_list_item"`
	ToDo       BlockCheckbox `json:"to_do"`
	Toggle     BlockText     `json:"toggle"`
	// TODO: child_page
}

type BlockCheckbox struct {
	BlockText
	Checked bool `json:"checked"`
}

type BlockText struct {
	Text []*RichText `json:"text"`
}

func (b *Block) GetChildren() (*PageContent, error) {
	if !b.HasChildren {
		return nil, fmt.Errorf("block does not have children")
	}

	return client.GetBlockChildren(b.ID)
}

func (c *Client) GetBlockChildren(blockID string) (*PageContent, error) {
	var pc PageContent
	err := c.doRequest(http.MethodGet, fmt.Sprintf("blocks/%s/children", blockID), nil, &pc)
	if err != nil {
		return nil, err
	}
	return &pc, nil
}

func (b *Block) AppendChildren(children []*Block) (*PageContent, error) {
	return client.AppendBlockChildren(b.ID, children)
}

func (c *Client) AppendBlockChildren(blockID string, children []*Block) (*PageContent, error) {
	body := map[string]interface{}{
		"children": children,
	}
	var pc PageContent
	err := c.doRequest(http.MethodPatch, fmt.Sprintf("blocks/%s/children", blockID), body, &pc)
	if err != nil {
		return nil, err
	}
	return &pc, nil
}
