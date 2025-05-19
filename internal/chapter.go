package internal

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	NumberStr string
	Number    int64
	ImgUrl    string
}

func GetChapterPages(chapters []*Chapter) {
	for _, chapter := range chapters {
		data := extractChapterPagesInfo(chapter.Link)
		chapter.Pages = &data
	}
}

// TODO use the same headless driver for all pages
func extractChapterPagesInfo(link string) []*Page {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // Set to true for production
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Run tasks in Chrome
	var results []map[string]interface{}
	err := chromedp.Run(ctx,
		// Navigate to the local HTML file
		chromedp.Navigate(link),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('div[id^="page"]'))
				.filter(div => /^page\d+$/.test(div.id))
				.map(div => ({
				id: div.id,
				dataSrc: div.querySelector('img[data-src]')?.getAttribute('data-src') || null
			}))
		`, &results),
	)
	if err != nil {
		log.Fatalf("Error running Chromedp: %v", err)
	}

	pages := make([]*Page, len(results))
	for idx, result := range results {
		page := &Page{}
		idStr := strings.Replace(result["id"].(string), "page", "", 1)
		page.NumberStr = result["id"].(string)
		page.Number, _ = strconv.ParseInt(idStr, 10, 64)
		page.ImgUrl = result["dataSrc"].(string)
		pages[idx] = page
	}

	return pages
}
