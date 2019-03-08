package filter

import (
	"github.com/selesy/go-html/node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Criteria func(*html.Node) bool
type Criterion Criteria

//Apply executes each of the criteria
func Apply(n *html.Node, criterion ...Criteria) *html.Node {
	if Test(n, criterion...) {
		return n
	}
	return nil
}

func Test(n *html.Node, criterion ...Criteria) bool {
	for _, c := range criterion {
		if !c(n) {
			return false
		}
	}
	return true
}

//Class is a convenience method that returns a FilterTest indicating
//whether or not an HTML node contains a class attribute with a value
//that matches the provided string.
func Class(c string) Criteria {
	return Matches(atom.Class, c)
}

// func Composite(criterion ...Criteria) {
// 	return func()
// }

//Exists returns a FilterTest that indicates whether or not a specific
//attribute is present on an HTML node.
func Exists(key atom.Atom) Criteria {
	return func(n *html.Node) bool {
		_, ok := node.Attr(n, key)
		return ok
	}
}

//Id is a convenience method that returns a FilterTest indicating whether
//or not an HTML node contains an id attribute with a value that matches
//the provided string.
func Id(i string) Criteria {
	return Matches(atom.Id, i)
}

//Matches returns a FilterTest that indicates whether an HTML node
//contains a matching attribute key/value pair.
func Matches(key atom.Atom, val string) Criteria {
	return func(n *html.Node) bool {
		v, ok := node.Attr(n, key)
		if ok && v == val {
			return true
		}
		return false
	}
}

//Tag returns a FilterTest that indicates whether an HTML node has
//a tag matching the provided string.
func Tag(t atom.Atom) Criteria {
	return func(n *html.Node) bool {
		return n.DataAtom == t
	}
}

//Type returns a FilterTest that indicates whether and HTML node is
//of the type passed as the sole parameter.
func Type(t html.NodeType) Criteria {
	return func(n *html.Node) bool {
		return n.Type == t
	}
}
