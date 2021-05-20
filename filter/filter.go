package filter

type Filter interface {
	IsValid() bool
	// fmt.Stringer
}
