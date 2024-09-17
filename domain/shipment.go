package domain

import (
	"time"
)

type ShippingLabelInterface interface {
	GeneratingPDF(size string) ([]byte, error)
}

type DataShipment struct {
	ShipmentID              int       `json:"shipment_id"`
	RefNo                   string    `json:"ref_no"`
	Awb                     string    `json:"awb"`
	CourierID               int       `json:"courier_id"`
	CourierServiceID        int       `json:"courier_service_id"`
	CourierName             string    `json:"courier_name"`
	CourierServiceName      string    `json:"courier_service_name"`
	CourierLogo             string    `json:"courier_logo"`
	AwbGeneratedDate        time.Time `json:"awb_generated_date"`
	OriginName              string    `json:"origin_name"`
	OriginAreaID            string    `json:"origin_area_id"`
	OriginPhone             string    `json:"origin_phone"`
	ShippingInsurance       int       `json:"shipping_insurance"`
	Price                   int       `json:"price"`
	DestinationName         string    `json:"destination_name"`
	DestinationPhone        string    `json:"destination_phone"`
	DestinationAreaID       string    `json:"destination_area_id"`
	DestinationAddress      string    `json:"destination_address"`
	DestinationZipcode      string    `json:"destination_zipcode"`
	Items                   []Items   `json:"items"`
	IsCod                   bool      `json:"is_cod"`
	CodFee                  int       `json:"cod_fee"`
	OriginDistrictCode      string    `json:"origin_district_code"`
	DestinationDistrictCode string    `json:"destination_district_code"`
	SortCode                string    `json:"sort_code"`
}

type Items struct {
	ItemName string      `json:"item_name"`
	ItemCode interface{} `json:"item_code"`
	Weight   int         `json:"weight"`
	Quantity int         `json:"quantity"`
	Value    int         `json:"value"`
}
