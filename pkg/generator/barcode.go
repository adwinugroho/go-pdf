package generator

import (
	"bytes"
	"image/png"
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(enc string, width, height int) ([]byte, error) {
	// Create the barcode
	codabar, err := code128.Encode(enc)
	if err != nil {
		log.Println("Error while generating barcode:", err)
		return nil, err
	}

	// Scale the barcode to the desired size
	resCode, err := barcode.Scale(codabar, width, height)
	if err != nil {
		log.Println("Error while scaling barcode:", err)
		return nil, err
	}

	// Create a buffer to write the PNG into memory
	var buf bytes.Buffer

	// Encode the barcode as PNG and write it to the buffer
	err = png.Encode(&buf, resCode)
	if err != nil {
		log.Println("Error while encoding PNG:", err)
		return nil, err
	}

	// Return the bytes from the buffer
	return buf.Bytes(), nil
}
