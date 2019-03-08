package scrape

import (
	"github.com/selesy/go-html/filter"
	yhat "github.com/yhat/scrape"
	"golang.org/x/net/html"
)

//composite wraps a criterion into a yhat Matcher
func composite(criterion ...filter.Criteria) yhat.Matcher {
	return func(n *html.Node) bool {
		return filter.Test(n, criterion...)
	}
}

func Find(node *html.Node, criterion ...filter.Criteria) (*html.Node, bool) {
	return yhat.Find(node, composite(criterion...))
}

func FindAll(node *html.Node, criterion ...filter.Criteria) []*html.Node {
	return yhat.FindAll(node, composite(criterion...))
}
