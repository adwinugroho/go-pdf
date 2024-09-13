package domain

type Invoice struct {
	ID          string         `json:"id"`
	OrderID     string         `json:"order_id"`
	Name        string         `json:"name"`
	Ship        Ship           `json:"ship"`
	Items       []ItemsInvoice `json:"items"`
	Status      string         `json:"status"`
	TotalAmount float64        `json:"total_amount"`
	Date        string         `json:"date"`
}

type Ship struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Country string `json:"country"`
}

type ItemsInvoice struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type InvoiceInterface interface {
	GeneratingPDF(filterDate []string) ([]byte, error)
}
