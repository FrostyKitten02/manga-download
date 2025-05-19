package internal

import "golang.org/x/net/html"

func findImgContainer(n *html.Node) *html.Node {
	return findNodeByClass(n, "uk-grid uk-grid-collapse")
}

func findChapterTable(n *html.Node) *html.Node {
	return findNodeByClass(n, "chapters")
}

func findNodeByClass(n *html.Node, class string) *html.Node {
	attributes := n.Attr
	if attributes != nil {
		for _, attr := range attributes {
			if attr.Key == "class" {
				if attr.Val == class {
					return n
				}
				break
			}
		}
	}

	childIter := n.ChildNodes()
	for child := range childIter {
		found := findNodeByClass(child, class)
		if found != nil {
			return found
		}
	}

	return nil
}
