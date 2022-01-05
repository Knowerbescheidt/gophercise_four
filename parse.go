package link

import (
	"fmt"
	"io"
	"strings"

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
		l := buildLink(node)
		links = append(links, l)
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
	}
	text := extractText(n)
	ret.Text = prepText(text)
	return ret
}

func prepText(text string) string {
	s := strings.Fields(text)
	t := strings.Join(s, " ")
	return t
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c) + " "
	}
	return ret
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
