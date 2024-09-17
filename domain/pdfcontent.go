package domain

type PDFContent struct {
	CenterLogoURL      string `json:"center_logo_url"`
	RightLogoURL       string `json:"right_logo_url"`
	LeftLogoURL        string `json:"left_logo_url"`
	IsUseUnboxingGuide bool   `json:"is_use_unboxing_guide"`
	FirstBarcodeImage  string `json:"first_barcode_image"`
	SecondBarcodeImage string `json:"second_barcode_image"`

	PrintGenerateDate       string      `json:"print_generate_date"`
	ShippingType            string      `json:"shipping_type"`             // COD/Non COD
	FullyDestinationAddress string      `json:"fully_destination_address"` // destination_area_id, destination_address, & destination_zipcode
	TotalPrice              int         `json:"total_price"`               // total price shipping
	TotalWeight             int         `json:"total_weight"`              // weight of product * qty ++
	TotalQty                int         `json:"total_qty"`                 // total qty product
	RefNo                   string      `json:"ref_no"`
	Awb                     string      `json:"awb"`
	CourierName             string      `json:"courier_name"`
	CourierLogo             string      `json:"courier_logo"`
	AwbGeneratedDate        string      `json:"awb_generated_date"`
	OriginName              string      `json:"origin_name"`
	OriginPhone             string      `json:"origin_phone"`
	ShippingInsurance       int         `json:"shipping_insurance"`
	Price                   int         `json:"price"`
	DestinationName         string      `json:"destination_name"`
	DestinationPhone        string      `json:"destination_phone"`
	Items                   []Items     `json:"items"`
	SortCode                interface{} `json:"sort_code"`
}
