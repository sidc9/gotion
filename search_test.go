package gotion_test

import (
	"testing"

	"github.com/matryer/is"
)

func TestSearch_ByDatabaseTitle(t *testing.T) {
	is := is.New(t)
	c := getClient(t)

	title := "Todo1s"
	db, err := c.SearchDatabaseByTitle(title)

	is.NoErr(err)
	is.Equal(db, nil)
}
