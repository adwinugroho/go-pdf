package service

import (
	"encoding/json"
	"fmt"
	"go-pdf/domain"
	"go-pdf/pkg/helper"
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

	invoices := make([]domain.Invoice, 0)
	err = json.Unmarshal(jsonBytes, &invoices)
	if err != nil {
		log.Println("Error while unmarshall:", err)
		return nil, err
	}
	resTemp, err := template.ProcessTemplate("./dist/template/invoice.html", "invoice", invoices)
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

func (s *invoiceImpl) transformDataForGenerate(shipments []domain.DataShipment, setting domain.Setting) []domain.PDFContent {
	result := make([]domain.PDFContent, 0)

	if len(shipments) > 0 {
		for _, eachShipment := range shipments {
			centerLogoURL := ""
			if setting.IsUseStoreLogo {
				centerLogoURL = setting.LogoURL
			}

			rightLogoURL := ""
			if eachShipment.CourierID == 1 {
				rightLogoURL = "J&T"
			}

			// errGenFirstBarcode := generator.GenerateBarcode(eachShipment.Awb, "../../dist/barcode/barcode-1.png", 600, 200)
			// if errGenFirstBarcode != nil {
			// 	log.Println("Error while generate first barcode:", errGenFirstBarcode)
			// }

			// errGenSecond := generator.GenerateBarcode(eachShipment.RefNo, "../../dist/barcode/barcode-2.png", 600, 200)
			// if errGenSecond != nil {
			// 	log.Println("Error while generate second barcode:", errGenSecond)
			// }

			firstBarcodeImage := "../../dist/barcode/barcode-1.png"
			secondBarcodeImage := "../../dist/barcode/barcode-2.png"

			fullyDestinationAddress := fmt.Sprintf("%s, %s, %s", eachShipment.DestinationAreaID, eachShipment.DestinationAddress, eachShipment.DestinationZipcode)
			totalQty := len(eachShipment.Items)
			totalWeight := 0
			if len(eachShipment.Items) > 0 {
				for _, eachItem := range eachShipment.Items {
					totalWeight = totalWeight + (eachItem.Weight * eachItem.Quantity)
				}
			}
			totalPrice := eachShipment.Price + eachShipment.ShippingInsurance
			awbDateString := eachShipment.AwbGeneratedDate.Format("02/01/2006 15:04")
			nowString := helper.TimeHostNow("Asia/Jakarta").Format("02/01/2006 15:04")
			shippingType := "Non COD"
			if eachShipment.IsCod {
				shippingType = "COD"
			}

			var dataPDF = domain.PDFContent{
				Awb:                eachShipment.Awb,
				RefNo:              eachShipment.RefNo,
				CourierName:        eachShipment.CourierName,
				OriginName:         eachShipment.OriginName,
				OriginPhone:        eachShipment.OriginPhone,
				ShippingInsurance:  eachShipment.ShippingInsurance,
				DestinationName:    eachShipment.DestinationName,
				DestinationPhone:   eachShipment.DestinationPhone,
				SortCode:           eachShipment.SortCode,
				Items:              eachShipment.Items,
				Price:              eachShipment.Price,
				IsUseUnboxingGuide: setting.IsUseUnboxingGuide,

				PrintGenerateDate:       nowString,
				FullyDestinationAddress: fullyDestinationAddress,
				TotalPrice:              totalPrice,
				TotalWeight:             totalWeight,
				TotalQty:                totalQty,
				AwbGeneratedDate:        awbDateString,
				ShippingType:            shippingType,

				CenterLogoURL:      centerLogoURL,
				RightLogoURL:       rightLogoURL,
				FirstBarcodeImage:  firstBarcodeImage,
				SecondBarcodeImage: secondBarcodeImage,
			}
			result = append(result, dataPDF)
		}
	}

	return result
}
