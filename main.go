package main

import (
	"fmt"
	"log"
	"manga-download/internal"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	c := &internal.Config{
		Manga: "https://mangakatana.com/manga/monster.8819",
	}
	chapters := internal.ExtractChapters(c)
	if chapters == nil {
		fmt.Println("No chapters found")
		return
	}
	internal.GetChapterPages(chapters)

	log.Print("FInished!")
}
