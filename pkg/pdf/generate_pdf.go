package pdf

import (
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateWithWKHtmlToPDF(html, pageSize string, isBW bool) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Printf("Error while creating PDF generator: %+v\n", err)
		return nil, err
	}

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set("Portrait")

	// Set the standard page size like "A4", "A5", "Letter", etc.
	pdfg.PageSize.Set(pageSize)

	if isBW {
		pdfg.Grayscale.Set(true)
	}

	// Create a new page from HTML content
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	page.Zoom.Set(1.0) // Adjust zoom if necessary to fit content properly
	pdfg.AddPage(page)

	// Generate the PDF
	err = pdfg.Create()
	if err != nil {
		log.Printf("Error while creating PDF: %+v\n", err)
		return nil, err
	}

	return pdfg.Bytes(), nil
}
