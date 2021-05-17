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

		t.Run("set lessThanOrEqual", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").LessThanOrEqual(2)
			is.Equal(*nf.Number.LessThanOrEqual, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set greaterThan", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").GreaterThan(2)
			is.Equal(*nf.Number.GreaterThan, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set lessThan", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").LessThan(2)
			is.Equal(*nf.Number.LessThan, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set Equals", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").Equals(2)
			is.Equal(*nf.Number.Equals, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set NotEqual", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").DoesNotEqual(2)
			is.Equal(*nf.Number.DoesNotEqual, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set IsEmpty", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").IsEmpty()
			is.Equal(*nf.Number.IsEmpty, true)
			is.Equal(nf.Property, "age")
		})

		t.Run("set IsNotEmpty", func(t *testing.T) {
			nf := gotion.NewNumberFilter("age").IsNotEmpty()
			is.Equal(*nf.Number.IsNotEmpty, true)
			is.Equal(nf.Property, "age")
		})
	})
}
