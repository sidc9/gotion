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

	Heading1  BlockText `json:"heading_1"`
	Heading2  BlockText `json:"heading_2"`
	Heading3  BlockText `json:"heading_3"`
	Paragraph BlockText `json:"paragraph"`
	Bullet    BlockText `json:"bulleted_list_item"`
	// TODO: more types of content
}

type BlockText struct {
	Text []*RichText `json:"text"`
}

func (b *Block) GetChildren() (*PageContent, error) {
	if !b.HasChildren {
		return nil, fmt.Errorf("block does not have children")
	}

	var pc PageContent
	err := client.doRequest(http.MethodGet, fmt.Sprintf("blocks/%s/children", b.ID), nil, &pc)
	if err != nil {
		return nil, err
	}
	return &pc, nil
}
