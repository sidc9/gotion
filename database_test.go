package main

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/matryer/is"
)

func TestListDatabases(t *testing.T) {
	is := is.New(t)

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	c := NewClient(apiKey)
	resp, err := c.ListDatabases()

	is.NoErr(err)
	is.Equal(resp.Object, "list")
	is.Equal(resp.HasMore, false)
	is.Equal(resp.NextCursor, "")

	is.Equal(len(resp.Results), 1)

	item := resp.Results[0]
	is.Equal(item.Object, "database")
	pretty.Println(item.Properties)

}

func TestGetDatabase(t *testing.T) {
	is := is.New(t)

	c := &Client{}
	t.Run("returns error if id is empty", func(t *testing.T) {
		db, err := c.GetDatabase("")
		is.Equal(err, fmt.Errorf("id is required"))
		is.Equal(db, nil)
	})

	db, err := c.GetDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.NoErr(err)
	is.Equal(db.Object, "database")
	is.Equal(db.ID, "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")

	pretty.Println(db)
}
