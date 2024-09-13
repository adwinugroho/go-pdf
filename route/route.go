package route

import (
	"go-pdf/route/handler"
	service "go-pdf/service/invoice"
	"time"

	"github.com/gofiber/fiber/v2"
)

func InvoiceRoute(app *fiber.App, timecontext time.Duration) {

	// Service
	invService := service.NewServiceInvoice()
	// handler
	invHandler := handler.NewInvoiceHandler(invService)
	// Route
	invRoute := app.Group("/invoice")
	invRoute.Post("/generate", invHandler.GeneratePDF)

}
