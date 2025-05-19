package main

import (
	"fmt"
	"log"
	"manga-download/internal"
)

func main() {
	c := &internal.Config{
		//Manga: "https://mangakatana.com/manga/monster.8819",
		Manga: "https://mangakatana.com/manga/akogare-no-onee-san.18122",
	}
	chapters := internal.ExtractChapters(c)
	if chapters == nil {
		fmt.Println("No chapters found")
		return
	}

	internal.GetChapterPages(chapters)

	log.Print("Finished!")
}
