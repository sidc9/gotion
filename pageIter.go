package gotion

type PageIter struct {
	client     *Client
	hasNext    bool
	nextCursor string
	index      int
	pages      []*Page
}

func (pi *PageIter) HasNext() bool {
	// TODO temporary to fix a panic. Write this properly.
	if pi.index == len(pi.pages) {
		return false
	}
	return pi.hasNext
}
func (pi *PageIter) GetNext() *Page {
	if pi.index < len(pi.pages) {
		pg := pi.pages[pi.index]
		pi.index++
		return pg
	}

	// make new query
	return nil
}
