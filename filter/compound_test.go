package filter

import (
	"encoding/json"
	"testing"

	"github.com/matryer/is"
)

func TestOrFilter(t *testing.T) {
	is := is.New(t)

	n := NewNumberFilter("p1").Equals(3)
	c := NewCheckboxFilter("p2").Equals(true)
	or, err := NewOrFilter(n, c)
	is.NoErr(err)

	is.Equal(or.Or[0], n)
	is.Equal(or.Or[1], c)

	b, err := json.Marshal(or)
	is.NoErr(err)

	want := `{"or":[{"number":{"equals":3},"property":"p1"},{"checkbox":{"equals":true},"property":"p2"}]}`
	is.Equal(want, string(b))
}

func TestAndFilter(t *testing.T) {
	is := is.New(t)

	n := NewNumberFilter("p1").Equals(3)
	c := NewCheckboxFilter("p2").Equals(true)
	and, err := NewAndFilter(n, c)
	is.NoErr(err)

	is.Equal(and.And[0], n)
	is.Equal(and.And[1], c)

	b, err := json.Marshal(and)
	is.NoErr(err)

	want := `{"and":[{"number":{"equals":3},"property":"p1"},{"checkbox":{"equals":true},"property":"p2"}]}`
	is.Equal(want, string(b))
}
