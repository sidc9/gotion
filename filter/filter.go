package filter

import "fmt"

type Filter interface {
	IsValid() bool
	fmt.Stringer
}
