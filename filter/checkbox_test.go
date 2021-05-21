package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestCheckboxFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Equals", func(t *testing.T) {
		cf := filter.NewCheckboxFilter("is_checked").Equals(false)
		is.Equal(cf.Property(), "is_checked")
		is.Equal(cf.Condition(), "equals")

		checkJSON(t, cf, `{"checkbox":{"equals":false},"property":"is_checked"}`)
	})

	t.Run("set DoesNotEqual", func(t *testing.T) {
		cf := filter.NewCheckboxFilter("is_checked").DoesNotEqual(true)
		is.Equal(cf.Condition(), "does_not_equal")
		is.Equal(cf.Property(), "is_checked")
	})
}
