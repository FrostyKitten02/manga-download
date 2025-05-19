package main

import (
	"fmt"
	"log"
	"manga-download/internal"
)

// TODO findout why we get wierd data in structs, but data is ok when its processing and in processing results, results also seem fine but chapters fomr extract chapters get messed up why??
func main() {
	c := &internal.Config{
		Manga: "https://mangakatana.com/manga/monster.8819",
	}
	chapters := internal.ExtractChapters(c)
	if chapters == nil {
		fmt.Println("No chapters found")
		return
	}
	subarr := (*(chapters))[0:2]
	internal.GetChapterPages(&subarr)

	log.Print("Finished!")
}
