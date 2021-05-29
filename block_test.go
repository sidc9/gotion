package gotion_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func setResponse(t *testing.T, c *gotion.Client, responseFile string) {
	filename := filepath.Join("testdata", responseFile)
	if saveResponse {
		t.Logf("    * saving to: %s\n", filename)
		c.SaveResponse(filename)
	} else {
		f, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		c.WithHTTPClient(&http.Client{
			Transport: &mockRoundTripper{
				fn: func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBuffer(f)),
					}, nil
				},
			},
		})
	}
}

func TestGetBlockChildren(t *testing.T) {
	is := is.New(t)

	c := getClient(t)
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
