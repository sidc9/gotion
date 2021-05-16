package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestFilter(t *testing.T) {
	is := is.New(t)

	t.Run("number filter", func(t *testing.T) {
		t.Run("set greaterThanOrEqual", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThanOrEqual(2)
			is.Equal(*nf.Number.GreaterThanOrEqual, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set greaterThan", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThan(2)
			is.Equal(*nf.Number.GreaterThan, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set Equals", func(t *testing.T) {
			nf := NewNumberFilter("age").Equals(2)
			is.Equal(*nf.Number.Equals, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("does not set greaterThanOrEqual if something else is already set", func(t *testing.T) {
			nf := NewNumberFilter("age").Equals(5)
			nf.GreaterThanOrEqual(2)
			is.Equal(*nf.Number.Equals, 5)
		})

		t.Run("does not set greaterThan if something else is already set", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThanOrEqual(5)
			nf.GreaterThan(2)
			is.Equal(*nf.Number.GreaterThanOrEqual, 5)
		})

		t.Run("does not set Equals if something else is already set", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThan(5)
			nf.Equals(2)
			is.Equal(*nf.Number.GreaterThan, 5)
		})
	})
}
