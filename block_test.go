package gotion_test

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func TestGetBlockChildren(t *testing.T) {
	is := is.New(t)

	respOut := filepath.Join("testdata", "get_block_children.txt")
	var req *http.Request
	setup(t, respOut, req)

	b := &gotion.Block{
		ID:          "d684b6f5-a784-4c06-836e-e596a764b404",
		HasChildren: true,
	}
	child, err := b.GetChildren()
	is.NoErr(err)

	is.Equal("yes he does", child.Results[0].Bullet.Text[0].PlainText)
}
