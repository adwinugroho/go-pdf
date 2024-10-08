package handler

import (
	"fmt"
	"go-pdf/domain"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type InvoiceHandler struct {
	invoiceService domain.InvoiceInterface
}

func NewInvoiceHandler(invService domain.InvoiceInterface) InvoiceHandler {
	return InvoiceHandler{invoiceService: invService}
}

func (h *InvoiceHandler) GeneratePDF(c *fiber.Ctx) error {
	var body domain.GeneratePDFReq
	err := c.BodyParser(&body)
	if err != nil {
		log.Println("error!", err)
		return c.Status(500).JSON(map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	if body.Filename == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"message": "Filename is required!",
		})
	}

	fileBytes, err := h.invoiceService.GeneratingPDF(body.Date)
	if err != nil {
		log.Println("[GeneratingPDF] error!", err)
		return c.Status(500).JSON(map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	// Ensure the directory exists
	dir := "/dist/download"
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Println("[MkdirAll] error!", err)
		return c.Status(500).JSON(map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	// Write the file
	err = os.WriteFile(fmt.Sprintf("%s/%s.pdf", dir, body.Filename), fileBytes, 0644)
	if err != nil {
		log.Println("[WriteFile] error!", err)
		return c.Status(500).JSON(map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"message": "successfully generating pdf!",
	})
}
