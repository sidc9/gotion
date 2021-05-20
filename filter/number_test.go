package filter_test

import (
	"encoding/json"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestNumberFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set greaterThanOrEqual", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").GreaterThanOrEqual(2)
		is.Equal(nf.Condition(), "greater_than_or_equal_to")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"greater_than_or_equal_to":2},"property":"age"}`)
	})

	t.Run("set lessThanOrEqual", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").LessThanOrEqual(2)
		is.Equal(nf.Condition(), "less_than_or_equal_to")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"less_than_or_equal_to":2},"property":"age"}`)
	})

	t.Run("set greaterThan", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").GreaterThan(2)
		is.Equal(nf.Condition(), "greater_than")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"greater_than":2},"property":"age"}`)
	})

	t.Run("set lessThan", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").LessThan(2)
		is.Equal(nf.Condition(), "less_than")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"less_than":2},"property":"age"}`)
	})

	t.Run("set Equals", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").Equals(2)
		is.Equal(nf.Condition(), "equals")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"equals":2},"property":"age"}`)
	})

	t.Run("set NotEqual", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").DoesNotEqual(2)
		is.Equal(nf.Condition(), "does_not_equal")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"does_not_equal":2},"property":"age"}`)
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").IsEmpty()
		is.Equal(nf.Condition(), "is_empty")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"is_empty":true},"property":"age"}`)
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		nf := filter.NewNumberFilter("age").IsNotEmpty()
		is.Equal(nf.Condition(), "is_not_empty")
		is.Equal(nf.Property(), "age")

		checkJSON(t, nf, `{"number":{"is_not_empty":true},"property":"age"}`)
	})

	t.Run("stringer", func(t *testing.T) {
		t.Run("equals", func(t *testing.T) {
			nf := filter.NewNumberFilter("age").Equals(2)
			want := "NumberFilter:age, equals:2"
			is.Equal(want, nf.String())
		})
	})
}

func checkJSON(t *testing.T, val interface{}, want string) {
	t.Helper()
	is := is.NewRelaxed(t)
	b, err := json.Marshal(val)
	is.NoErr(err)
	is.Equal(want, string(b))
}
