package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestMultiMultiSelectFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Contains", func(t *testing.T) {
		f := filter.NewMultiSelectFilter("dummy").Contains("something")
		is.Equal(f.Property(), "dummy")
		is.Equal(f.Condition(), "contains")
		checkJSON(t, f, `{"multi_select":{"contains":"something"},"property":"dummy"}`)
	})

	t.Run("set DoesNotContain", func(t *testing.T) {
		f := filter.NewMultiSelectFilter("dummy").DoesNotContain("something")
		is.Equal(f.Property(), "dummy")
		is.Equal(f.Condition(), "does_not_contain")
		checkJSON(t, f, `{"multi_select":{"does_not_contain":"something"},"property":"dummy"}`)
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		f := filter.NewMultiSelectFilter("age").IsEmpty()
		is.Equal(f.Condition(), "is_empty")
		is.Equal(f.Property(), "age")

		checkJSON(t, f, `{"multi_select":{"is_empty":true},"property":"age"}`)
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		f := filter.NewMultiSelectFilter("age").IsNotEmpty()
		is.Equal(f.Condition(), "is_not_empty")
		is.Equal(f.Property(), "age")

		checkJSON(t, f, `{"multi_select":{"is_not_empty":true},"property":"age"}`)
	})
}
