package internal

import (
	"golang.org/x/net/html"
	"log"
)

type Chapter struct {
	Link    string
	Chapter string
	Pages   *[]*Page
}

func ExtractChapters(c *Config) *[]*Chapter {
	node := download(c.Manga)
	if node == nil {
		log.Println("No chapters found")
		return nil
	}
	chapterTable := findChapterTable(node)
	chapters := extractChapters(chapterTable)

	if chapters == nil || len(chapters) == 0 {
		return nil
	}

	return &chapters
}

func extractChapters(chapters *html.Node) []*Chapter {
	if chapters == nil {
		return nil
	}
	c1 := chapters.FirstChild
	if c1 == nil {
		return nil
	}

	c2 := c1.FirstChild
	if c2 == nil {
		return nil
	}

	//TODO find number of chapters and pre-allocate
	chaptersData := make([]*Chapter, 0)
	childIter := c2.ChildNodes()
	for child := range childIter {
		if child.Data != "tr" {
			continue
		}
		w1 := child.FirstChild
		if w1 == nil {
			log.Println("Error getting chapter data, w1")
			continue
		}

		w2 := w1.FirstChild
		if w2 == nil {
			log.Println("Error getting chapter data, w2")
			continue
		}

		aTag := w2.FirstChild
		if aTag == nil {
			log.Println("Error getting chapter data, aTag")
			continue
		}

		chapter := &Chapter{}
		atrributes := aTag.Attr
		for _, attr := range atrributes {
			if attr.Key == "href" {
				chapter.Link = attr.Val
			}

			chapter.Chapter = aTag.FirstChild.Data
		}
		chaptersData = append(chaptersData, chapter)
	}

	return chaptersData
}
