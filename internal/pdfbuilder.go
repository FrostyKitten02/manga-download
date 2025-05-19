package internal

import (
	"github.com/jung-kurt/gofpdf"
	"image"
	"io"
	"log"
	"net/http"
	"os"
)

func CreatePdf(chapter Chapter) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, page := range chapter.Pages {
		file, fileErr := downloadImg(page.ImgUrl)
		if fileErr != nil {
			log.Fatal(fileErr)
			return
		}
		pdf.AddPage()
		pdf.Image(file.Name(), 0, 0, 0, 0, false, "", 0, "")
	}

	err := pdf.OutputFileAndClose("./test.pdf")
	if err != nil {
		log.Fatal(err)
		return
	}
}

// TODO fix 0 bytes in image, probably because of image.Decode??
func downloadImg(url string) (*os.File, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, format, imgErr := image.Decode(res.Body)
	if imgErr != nil {
		return nil, imgErr
	}

	imgFile, fileErr := os.CreateTemp("", "*."+format)
	if fileErr != nil {
		return nil, fileErr
	}
	defer imgFile.Close()
	_, writeErr := io.Copy(imgFile, res.Body)
	if writeErr != nil {
		return nil, writeErr
	}

	return imgFile, nil
}
