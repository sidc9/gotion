package gotion_test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func TestGetPage(t *testing.T) {
	is := is.New(t)

	t.Run("returns error if id is empty", func(t *testing.T) {
		c := &gotion.Client{}
		db, err := c.GetPage("")
		is.Equal(err, fmt.Errorf("id is required"))
		is.Equal(db, nil)
	})

	respOut := filepath.Join("testdata", "get_page.txt")
	var req *http.Request
	c := setup(t, respOut, req)

	pageID := "a0e3feca-85c9-440f-91cc-8c367d6aa9f4"
	page, err := c.GetPage(pageID)
	if err != nil {
		fmt.Println(err)
	}

	is.Equal(page.Object, "page")
	is.Equal(page.ID, "a0e3feca-85c9-440f-91cc-8c367d6aa9f4")
	is.Equal(page.Parent["database_id"], "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.Equal(page.ParentID(), "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.Equal(page.Title(), "john")
}
