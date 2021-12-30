package link

import (
	"io"
)

//Link represents a link in an html document.
type Link struct {
	Href string
	Text string
}

//Parse will take in an html document and will return a slice of link parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
