package internal

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Page struct {
	NumberStr string
	Number    int64
	ImgUrl    string
}

// TODO use the same headless driver for all pages
func extractChapterPagesInfo(link string) {
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
			Array.from(document.querySelectorAll('div[id^="page"]')).map(div => ({
				id: div.id,
				dataSrc: div.querySelector('img[data-src]')?.getAttribute('data-src') || null
			}))
		`, &results),
	)
	if err != nil {
		log.Fatalf("Error running Chromedp: %v", err)
	}
}
