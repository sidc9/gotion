package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestDateFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Equals", func(t *testing.T) {
		df := filter.NewDateFilter("due_date").Equals("hello")
		is.Equal(df.Condition(), "equals")
		is.Equal(df.Property(), "due_date")

		checkJSON(t, df, `{"date":{"equals":"hello"},"property":"due_date"}`)
	})

	t.Run("set Before", func(t *testing.T) {
		df := filter.NewDateFilter("due_date").Before("hello")
		is.Equal(df.Condition(), "before")
		is.Equal(df.Property(), "due_date")

		checkJSON(t, df, `{"date":{"before":"hello"},"property":"due_date"}`)
	})

	t.Run("set After", func(t *testing.T) {
		df := filter.NewDateFilter("due_date").After("hello")
		is.Equal(df.Condition(), "after")
		is.Equal(df.Property(), "due_date")

		checkJSON(t, df, `{"date":{"after":"hello"},"property":"due_date"}`)
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		df := filter.NewDateFilter("due_date").IsEmpty()
		is.Equal(df.Condition(), "is_empty")
		is.Equal(df.Property(), "due_date")

		checkJSON(t, df, `{"date":{"is_empty":true},"property":"due_date"}`)
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		df := filter.NewDateFilter("due_date").IsNotEmpty()
		is.Equal(df.Condition(), "is_not_empty")
		is.Equal(df.Property(), "due_date")

		checkJSON(t, df, `{"date":{"is_not_empty":true},"property":"due_date"}`)
	})
}
