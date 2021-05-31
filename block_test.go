package gotion_test

import (
	"net/http"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func TestGetBlockChildren(t *testing.T) {
	is := is.New(t)

	c := getClient(t)
	setResponse(t, c, "get_block_children.txt", http.MethodGet, "/v1/blocks/7eb342cf-c40c-43b3-b2ee-c0f8c2c79d88/children")

	b := &gotion.Block{
		ID:          "7eb342cf-c40c-43b3-b2ee-c0f8c2c79d88",
		HasChildren: true,
	}

	child, err := b.GetChildren()
	is.NoErr(err)

	is.Equal(4, len(child.Results))
	is.Equal("yes he does", child.Results[0].BulletList.Text[0].PlainText)
	is.Equal("first item", child.Results[1].NumberList.Text[0].PlainText)
	is.Equal("task to do", child.Results[2].ToDo.Text[0].PlainText)
	is.Equal(false, child.Results[2].ToDo.Checked)
	is.Equal("this is a toggle", child.Results[3].Toggle.Text[0].PlainText)
	is.Equal(true, child.Results[3].HasChildren)

	setResponse(t, c, "get_block_children_2.txt", http.MethodGet, "/v1/blocks/8dfd5961-783b-4707-8523-0de4a5dbd580/children")
	toggleChild, err := child.Results[3].GetChildren()
	is.NoErr(err)

	is.Equal("some text", toggleChild.Results[0].Paragraph.Text[0].PlainText)
}

func TestAppendBlockChildren(t *testing.T) {
	if saveResponse {
		t.Skip("blocks cannot be deleted, skipping test which modifies a page")
	}

	is := is.New(t)

	c := getClient(t)
	setResponse(t, c, "append_block_children.txt", http.MethodPatch, "/v1/blocks/63d396a6-9687-4ea6-80c8-eea4c6212658/children")

	b := &gotion.Block{
		ID: "63d396a6-9687-4ea6-80c8-eea4c6212658",
	}

	children := []*gotion.Block{
		{
			Object: "block",
			Type:   "heading_1",
			Heading1: gotion.BlockText{
				Text: []*gotion.RichText{
					{
						Type: "text",
						Text: gotion.Text{
							Content: "her name is Mary",
						},
					},
				},
			},
		},
	}

	_, err := b.AppendChildren(children)
	is.NoErr(err)
}
