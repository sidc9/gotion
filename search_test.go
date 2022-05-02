package gotion_test

import (
	"errors"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func TestSearch_ByDatabaseTitle(t *testing.T) {
	c := getClient(t)

	t.Run("successful search", func(t *testing.T) {
		is := is.New(t)
		title := "Todos"
		db, err := c.SearchDatabaseByTitle(title)

		is.NoErr(err)
		is.Equal(db, &gotion.Database{})
	})

	t.Run("not found", func(t *testing.T) {
		is := is.New(t)
		title := "Todo1s"
		db, err := c.SearchDatabaseByTitle(title)

		is.True(errors.Is(err, gotion.ErrNotFound))
		is.Equal(db, nil)
	})
}
