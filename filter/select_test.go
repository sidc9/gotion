package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestSelectFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Equals", func(t *testing.T) {
		cf := filter.NewSelectFilter("dummy").Equals("something")
		is.Equal(*cf.Select.Equals, "something")
		is.Equal(cf.Property, "dummy")
	})

	t.Run("set DoesNotEqual", func(t *testing.T) {
		cf := filter.NewSelectFilter("dummy").DoesNotEqual("something")
		is.Equal(*cf.Select.DoesNotEqual, "something")
		is.Equal(cf.Property, "dummy")
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		nf := filter.NewSelectFilter("age").IsEmpty()
		is.Equal(*nf.Select.IsEmpty, true)
		is.Equal(nf.Property, "age")
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		nf := filter.NewSelectFilter("age").IsNotEmpty()
		is.Equal(*nf.Select.IsNotEmpty, true)
		is.Equal(nf.Property, "age")
	})
}
