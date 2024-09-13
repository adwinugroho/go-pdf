package domain

type GeneratePDFReq struct {
	Date     []string `json:"date"`
	Filename string   `json:"filename"`
}
