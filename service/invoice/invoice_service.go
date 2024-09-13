package service

import (
	"encoding/json"
	"go-pdf/domain"
	"go-pdf/pkg/pdf"
	"go-pdf/pkg/template"
	"log"
	"os"
)

type invoiceImpl struct{}

func NewServiceInvoice() domain.InvoiceInterface {
	return &invoiceImpl{}
}

func (s *invoiceImpl) GeneratingPDF(filterDate []string) ([]byte, error) {
	jsonBytes, err := os.ReadFile("./dist/example/data.json")
	if err != nil {
		log.Println("Error while read dummy data:", err)
		return nil, err
	}

	// log.Printf("string jsonBytes:%+v\n", string(jsonBytes))

	invoices := make([]domain.Invoice, 0)
	err = json.Unmarshal(jsonBytes, &invoices)
	if err != nil {
		log.Println("Error while unmarshall:", err)
		return nil, err
	}

	if len(filterDate) > 0 {
		startDate := filterDate[0]
		endDate := filterDate[1]

		data := make([]domain.Invoice, 0)
		for _, eachInv := range invoices {
			if eachInv.Date >= startDate && eachInv.Date <= endDate {
				data = append(data, eachInv)
			}
		}

		resTemp, err := template.ProcessTemplate("./dist/template/invoice.html", data)
		if err != nil {
			log.Println("Error while process template:", err)
			return nil, err
		}

		resPDFBytes, err := pdf.GenerateWithWKHtmlToPDF(resTemp, "A5", false)
		if err != nil {
			log.Println("Error while generate PDF with WKhtml:", err)
			return nil, err
		}

		return resPDFBytes, nil
	} else {
		resTemp, err := template.ProcessTemplate("./dist/template/invoice.html", invoices)
		if err != nil {
			log.Println("Error while process template:", err)
			return nil, err
		}

		resPDFBytes, err := pdf.GenerateWithWKHtmlToPDF(resTemp, "A6", false)
		if err != nil {
			log.Println("Error while generate pdf with WKhtml:", err)
			return nil, err
		}
		return resPDFBytes, nil
	}
}
