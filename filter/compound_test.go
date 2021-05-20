package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestOrFilter(t *testing.T) {
	n := NewNumberFilter("p1")
	c := NewCheckboxFilter("p2")
	or, _ := NewOrFilter(n, c)

	is := is.New(t)
	is.Equal(or.Or[0], n)
	is.Equal(or.Or[1], c)

	b, err := json.Marshal(or)
	is.NoErr(err)
	fmt.Println(string(b))
}
