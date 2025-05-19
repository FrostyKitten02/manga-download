package internal

import (
	"golang.org/x/net/html"
	"log"
	"strconv"
	"strings"
)

type Page struct {
	NumberStr string
	Number    int64
	ImgUrl    string
}

func DownloadChapterImgs(c *Chapter) *string {
	doc := download(c.Link)
	if doc == nil {
		return nil
	}

	imgs := findImgContainer(doc)
	if imgs == nil {
		log.Println("Failed to find chapter images container")
		return nil
	}

	pages := getChapterPages(imgs)
	if pages == nil {
		return nil
	}

	return nil
}

func getChapterPages(imgWrapper *html.Node) []*Page {
	children := imgWrapper.ChildNodes()

	pages := make([]*Page, 0)
	for child := range children {
		//TODO better break out of this loop?
		skip := true
		for _, attr := range child.Attr {
			if attr.Key == "class" {
				skip = strings.Contains(attr.Val, "wrap_img uk-width-1-1")
				break
			}
		}

		if skip {
			continue
		}

		page := Page{}
		extractPageNumber(child, &page)
		extractImgUrl(child, &page)
		pages = append(pages, &page)
	}

	return pages
}

func extractImgUrl(parent *html.Node, page *Page) {
	node := parent.FirstChild
	attrs := node.Attr
	for _, attr := range attrs {
		if attr.Key == "data-src" {
			page.ImgUrl = attr.Val
		}
	}
}

func extractPageNumber(node *html.Node, page *Page) {
	attrs := node.Attr
	for _, attr := range attrs {
		if attr.Key == "id" {
			page.NumberStr = attr.Val
			numStr := strings.ReplaceAll(attr.Val, "page", "")
			page.Number, _ = strconv.ParseInt(numStr, 10, 64)
		}
	}
}
