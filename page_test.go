package gotion_test

import (
	"fmt"
	"net/http"
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

	c := getClient(t)
	setResponse(t, c, "get_page.txt", http.MethodGet, "/v1/pages/a0e3feca-85c9-440f-91cc-8c367d6aa9f4")

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
	is.Equal(false, page.Archived)
}

func TestGetPageContent(t *testing.T) {
	is := is.New(t)

	t.Run("returns error if id is empty", func(t *testing.T) {
		c := &gotion.Client{}
		db, err := c.GetPageContent("")
		is.Equal(err, fmt.Errorf("id is required"))
		is.Equal(db, nil)
	})

	c := getClient(t)
	setResponse(t, c, "get_page_content.txt", http.MethodGet, "/v1/blocks/a0e3feca-85c9-440f-91cc-8c367d6aa9f4/children")

	pageID := "a0e3feca-85c9-440f-91cc-8c367d6aa9f4"
	content, err := c.GetPageContent(pageID)
	is.NoErr(err)

	is.Equal("John Smith", content.Results[0].Heading1.Text[0].PlainText)
	is.Equal("He loves ", content.Results[1].Paragraph.Text[0].PlainText)
	is.Equal("reading ", content.Results[1].Paragraph.Text[1].PlainText)
	is.Equal("blue", content.Results[1].Paragraph.Text[1].Annotations.Color)
	is.Equal("and ", content.Results[1].Paragraph.Text[2].PlainText)
	is.Equal("cycling", content.Results[1].Paragraph.Text[3].PlainText)
	is.Equal(true, content.Results[1].Paragraph.Text[3].Annotations.Bold)
	is.Equal(".", content.Results[1].Paragraph.Text[4].PlainText)
}

func TestUpdatePageProperty(t *testing.T) {
	is := is.New(t)

	t.Run("returns error if id is empty", func(t *testing.T) {
		c := &gotion.Client{}
		page, err := c.UpdatePageProperty("", "", gotion.PageProperty{})
		is.Equal(err, fmt.Errorf("page id is required"))
		is.Equal(page, nil)
	})

	c := getClient(t)
	setResponse(t, c, "update_page_property.txt", http.MethodGet, "/v1/pages/a0e3feca-85c9-440f-91cc-8c367d6aa9f4")

	pageID := "a0e3feca-85c9-440f-91cc-8c367d6aa9f4"
	page, err := c.UpdatePageProperty(pageID, "age", gotion.PageProperty{Number: 21})
	is.NoErr(err)

	is.Equal(page.Object, "page")
	is.Equal(page.ID, "a0e3feca-85c9-440f-91cc-8c367d6aa9f4")
	is.Equal(page.Parent["database_id"], "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.Equal(page.ParentID(), "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.Equal(page.Title(), "john")
	is.Equal(page.Properties["age"].Number, 21)
	is.Equal(false, page.Archived)

	t.Cleanup(func() {
		_, err := c.UpdatePageProperty(pageID, "age", gotion.PageProperty{Number: 23})
		is.NoErr(err)
	})
}
