package generator

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(enc, filename string, width, heigth int) error {
	// Create the barcode
	codabar, err := code128.Encode(enc)
	if err != nil {
		log.Println("Error while generate barcode:", err)
		return err
	}

	resCode, err := barcode.Scale(codabar, width, heigth)
	if err != nil {
		log.Println("Error while scale codabar barcode:", err)
		return err
	}

	// Ensure the directory exists
	dir := "/dist/barcode"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Println("Error while mkdir barcode:", err)
		return err
	}

	// create the output file
	file, _ := os.Create(fmt.Sprintf("%s/%s", dir, filename))
	defer file.Close()

	// encode the barcode as png
	err = png.Encode(file, resCode)
	if err != nil {
		log.Println("Error encode png:", err)
		return err
	}

	return nil
}
