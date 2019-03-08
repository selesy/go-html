package node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//Attr is an html.Node helper function to retrieve an attribute's
//value from the node's attribute array.  It returns both the string
//value of the attribute and true if the attribute exists, otherwise
//it returns false.
func Attr(n *html.Node, key atom.Atom) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key.String() {
			return attr.Val, true
		}
	}
	return "", false
}
