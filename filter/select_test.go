package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestSelectFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Equals", func(t *testing.T) {
		sf := filter.NewSelectFilter("dummy").Equals("something")
		is.Equal(sf.Condition(), "equals")
		is.Equal(sf.Property(), "dummy")

		checkJSON(t, sf, `{"property":"dummy","select":{"equals":"something"}}`)
	})

	t.Run("set DoesNotEqual", func(t *testing.T) {
		sf := filter.NewSelectFilter("dummy").DoesNotEqual("something")
		is.Equal(sf.Condition(), "does_not_equal")
		is.Equal(sf.Property(), "dummy")

		checkJSON(t, sf, `{"property":"dummy","select":{"does_not_equal":"something"}}`)
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		sf := filter.NewSelectFilter("age").IsEmpty()
		is.Equal(sf.Condition(), "is_empty")
		is.Equal(sf.Property(), "age")

		checkJSON(t, sf, `{"property":"age","select":{"is_empty":true}}`)
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		sf := filter.NewSelectFilter("age").IsNotEmpty()
		is.Equal(sf.Condition(), "is_not_empty")
		is.Equal(sf.Property(), "age")

		checkJSON(t, sf, `{"property":"age","select":{"is_not_empty":true}}`)
	})
}
