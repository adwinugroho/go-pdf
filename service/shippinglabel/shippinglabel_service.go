package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-pdf/domain"
	"go-pdf/pkg/generator"
	"go-pdf/pkg/helper"
	"go-pdf/pkg/pdf"
	"go-pdf/pkg/template"
	"go-pdf/pkg/upload"
	"log"
	"os"
)

type shippingLabelImpl struct{}

func NewServiceShippingLabel() domain.ShippingLabelInterface {
	return &shippingLabelImpl{}
}

type rawData struct {
	Setting       domain.Setting        `json:"setting"`
	DataShipments []domain.DataShipment `json:"data"`
}

func (s *shippingLabelImpl) GeneratingPDF(size string) ([]byte, error) {
	jsonBytes, err := os.ReadFile("./dist/example/shipment.json")
	if err != nil {
		log.Println("Error while read dummy data:", err)
		return nil, err
	}

	var dummyData rawData
	err = json.Unmarshal(jsonBytes, &dummyData)
	if err != nil {
		log.Println("Error while unmarshall:", err)
		return nil, err
	}

	pdfData := s.transformDataForGenerate(dummyData.DataShipments, dummyData.Setting)

	resTemp, err := template.ProcessTemplate(fmt.Sprintf("./dist/template/template-%s.html", size), size, pdfData)
	if err != nil {
		log.Println("Error while process template:", err)
		return nil, err
	}

	resPDFBytes, err := pdf.GenerateWithWKHtmlToPDF(resTemp, size, false)
	if err != nil {
		log.Println("Error while generate pdf with WKhtml:", err)
		return nil, err
	}
	return resPDFBytes, nil
}

func (s *shippingLabelImpl) transformDataForGenerate(shipments []domain.DataShipment, setting domain.Setting) []domain.PDFContent {
	result := make([]domain.PDFContent, 0)

	if len(shipments) > 0 {
		for _, eachShipment := range shipments {
			centerLogoURL := ""
			if setting.IsUseStoreLogo {
				centerLogoURL = setting.LogoURL
			}

			rightLogoURL := eachShipment.CourierLogo
			leftLogoURL := "https://shipment.jubelio.com/img/favicon.svg"

			fBarcodeBytes, errGenFirstBarcode := generator.GenerateBarcode(eachShipment.Awb, 600, 100)
			if errGenFirstBarcode != nil {
				log.Println("Error while generate first barcode:", errGenFirstBarcode)
			}
			resFBarcode, err := upload.UploadImage(context.Background(), "barcode-1.png", fBarcodeBytes)
			if err != nil {
				log.Println("Error while generate first barcode:", errGenFirstBarcode)
			}

			sBarcodeBytes, errGenSecond := generator.GenerateBarcode(eachShipment.RefNo, 600, 100)
			if errGenSecond != nil {
				log.Println("Error while generate second barcode:", errGenSecond)
			}
			resSBarcode, err := upload.UploadImage(context.Background(), "barcode-2.png", sBarcodeBytes)
			if err != nil {
				log.Println("Error while generate first barcode:", errGenFirstBarcode)
			}

			firstBarcodeImage := resFBarcode.ObjectURL
			secondBarcodeImage := resSBarcode.ObjectURL

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
				CourierLogo:        eachShipment.CourierLogo,

				PrintGenerateDate:       nowString,
				FullyDestinationAddress: fullyDestinationAddress,
				TotalPrice:              totalPrice,
				TotalWeight:             totalWeight,
				TotalQty:                totalQty,
				AwbGeneratedDate:        awbDateString,
				ShippingType:            shippingType,

				CenterLogoURL:      centerLogoURL,
				RightLogoURL:       rightLogoURL,
				LeftLogoURL:        leftLogoURL,
				FirstBarcodeImage:  firstBarcodeImage,
				SecondBarcodeImage: secondBarcodeImage,
			}
			result = append(result, dataPDF)
		}
	}

	return result
}
