package gotion_test

import (
	"testing"
)

// GetNext()
// - must query API if cache is empty
// - must return latest element from cache (if available)
// - must modify the page-size in query so that the 'limit' is respected.
func TestPageIter(t *testing.T) {
}
