package internal

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
	"image/png"
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
		pageW, pageH := pdf.GetPageSize()
		pdf.Image(file.Name(), 0, 0, pageW, pageH, false, "", 0, "")
	}

	err := pdf.OutputFileAndClose("./" + chapter.Chapter + ".pdf")
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

	img, format, imgErr := image.Decode(res.Body)
	if imgErr != nil {
		return nil, imgErr
	}

	imgFile, fileErr := os.CreateTemp("", "*."+format)
	if fileErr != nil {
		return nil, fileErr
	}
	defer imgFile.Close()

	switch format {
	case "png":
		err = png.Encode(imgFile, img)
	case "jpeg", "jpg":
		err = jpeg.Encode(imgFile, img, &jpeg.Options{Quality: 100})
	default:
		imgFile.Close()
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	return imgFile, nil
}
