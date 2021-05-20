package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestOrFilter(t *testing.T) {
	is := is.New(t)

	n := NewNumberFilter("p1").Equals(3)
	// c := NewCheckboxFilter("p2")
	or, err := NewOrFilter(n)
	is.NoErr(err)

	is.Equal(or.Or[0], n)
	// is.Equal(or.Or[1], c)

	b, err := json.Marshal(or)
	is.NoErr(err)
	fmt.Println(string(b))
}
