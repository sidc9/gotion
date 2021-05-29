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
	// req := &http.Request{
	//     Method: http.MethodGet,
	//     URL: &url.URL{
	//         Path: "blocks/7eb342cf-c40c-43b3-b2ee-c0f8c2c79d88/children",
	//     },
	// }
	// setResponse(t, c, req, "get_block_children.txt")
	setResponse(t, c, "get_block_children.txt")

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

	// req.URL = &url.URL{
	//     Path: "blocks/8dfd5961-783b-4707-8523-0de4a5dbd580/children",
	// }
	// setResponse(t, c, req, "get_block_children_2.txt")
	setResponse(t, c, "get_block_children_2.txt")
	toggleChild, err := child.Results[3].GetChildren()
	is.NoErr(err)

	is.Equal("some text", toggleChild.Results[0].Paragraph.Text[0].PlainText)

}

type mockRoundTripper struct {
	fn func(*http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.fn(r)
}
