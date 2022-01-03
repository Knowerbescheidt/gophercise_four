package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

//Link represents a link in an html document.
type Link struct {
	Href string
	Text string
}

//Parse will take in an html document and will return a slice of link parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		fmt.Println(node)
		links = append(links, Link{node.Data, node.Type})
	}
	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
