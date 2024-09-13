package domain

type MessageBody struct {
	TenantID   string         `json:"tenant_id"`
	DownloadID string         `json:"download_id"`
	Setting    Setting        `json:"setting"`
	Data       []DataShipment `json:"data"`
}

type Setting struct {
	LogoURL            string `json:"logo_url"`
	IsUseSku           bool   `json:"is_use_sku"`
	PaperSize          string `json:"paper_size"`
	IsBlackWhite       bool   `json:"is_black_white"`
	IsUseStoreLogo     bool   `json:"is_use_store_logo"`
	IsUseOriginPhone   bool   `json:"is_use_origin_phone"`
	IsUseProductName   bool   `json:"is_use_product_name"`
	IsUseUnboxingGuide bool   `json:"is_use_unboxing_guide"`
}
