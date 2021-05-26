package gotion_test

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
)

func TestGetBlockChildren(t *testing.T) {
	is := is.New(t)

	respOut := filepath.Join("testdata", "get_block_children.txt")
	var req *http.Request
	c := setup(t, respOut, req)

	pageID := "a0e3feca-85c9-440f-91cc-8c367d6aa9f4"
	content, err := c.GetPageContent(pageID)
	is.NoErr(err)

	err = content.Results[1].GetChildren()
	is.NoErr(err)
}
