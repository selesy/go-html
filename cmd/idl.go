package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const moz = "https://dxr.mozilla.org/mozilla-central/source/dom/webidl/"
const w3c = "http://w3c.github.io/html/single-page.html"
const whatwg = "https://html.spec.whatwg.org/#the-body-element"

type MozIDL struct {
	Path        string
	Pescription string
	Date        string
	Size        string
}

func mozHTML() error {
	resp, err := http.Get(moz)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info(resp.StatusCode)

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	class := "folder-content"
	var f func(*html.Node) *html.Node
	f = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "table" && attribute(n, "class", class) {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			n := f(c)
			if n != nil {
				return n
			}
		}
		return nil
	}

	table := f(doc)

	if table == nil {
		return nil
	}
	log.Info(table.Data)

	// tbody, tr, td/td/td/td =

	return nil
}

func attribute(n *html.Node, attrKey string, attrVal string) bool {
	for _, attr := range n.Attr {
		if attr.Key == attrKey && attr.Val == attrVal {
			return true
		}
	}
	return false
}

type finder struct{}

type finderTest func(*html.Node) bool

func (f finder) passes(n *html.Node, tests ...finderTest) bool {
	for _, test := range tests {
		if !test(n) {
			return false
		}
	}
	return true
}
