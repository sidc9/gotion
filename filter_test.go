package gotion_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func TestFilter(t *testing.T) {
	is := is.New(t)

	t.Run("number filter", func(t *testing.T) {
		t.Run("set greaterThanOrEqual", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").GreaterThanOrEqual(2)
			is.Equal(*nf.Number.GreaterThanOrEqual, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set greaterThan", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").GreaterThan(2)
			is.Equal(*nf.Number.GreaterThan, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set Equals", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").Equals(2)
			is.Equal(*nf.Number.Equals, 2)
			is.Equal(nf.Property, "age")
		})
	})
}
