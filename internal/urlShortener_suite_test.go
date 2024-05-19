package internal

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"golang.org/x/net/html"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AgentTarget Reconciler Suite")
}

func fetchHref(n *html.Node) string {
	// Find the anchor tag
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hrefValue := fetchHref(c)
		if hrefValue != "" {
			return hrefValue
		}
	}
	return ""
}
	
