package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestMultiMultiSelectFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Contains", func(t *testing.T) {
		cf := filter.NewMultiSelectFilter("dummy").Contains("something")
		is.Equal(*cf.MultiSelect.Contains, "something")
		is.Equal(cf.Property, "dummy")
	})

	t.Run("set DoesNotContain", func(t *testing.T) {
		cf := filter.NewMultiSelectFilter("dummy").DoesNotContain("something")
		is.Equal(*cf.MultiSelect.DoesNotContain, "something")
		is.Equal(cf.Property, "dummy")
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		nf := filter.NewMultiSelectFilter("age").IsEmpty()
		is.Equal(*nf.MultiSelect.IsEmpty, true)
		is.Equal(nf.Property, "age")
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		nf := filter.NewMultiSelectFilter("age").IsNotEmpty()
		is.Equal(*nf.MultiSelect.IsNotEmpty, true)
		is.Equal(nf.Property, "age")
	})
}
