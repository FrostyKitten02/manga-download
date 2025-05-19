package internal

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func download(link string) *html.Node {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()

	node, parseErr := html.Parse(res.Body)
	if parseErr != nil {
		log.Fatal(parseErr)
		return nil
	}

	return node
}
