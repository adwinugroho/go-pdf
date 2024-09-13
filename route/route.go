package route

import (
	"go-pdf/route/handler"
	invService "go-pdf/service/invoice"
	slService "go-pdf/service/shippinglabel"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, timecontext time.Duration) {

	// Service
	invoiceService := invService.NewServiceInvoice()
	shippingLabelService := slService.NewServiceShippingLabel()
	// handler
	invHandler := handler.NewInvoiceHandler(invoiceService)
	slHandler := handler.NewShippingLabelHandler(shippingLabelService)

	// Route
	invRoute := app.Group("/invoice")
	invRoute.Post("/generate", invHandler.GeneratePDF)

	slRoute := app.Group("/shipping-label")
	slRoute.Post("generate", slHandler.GeneratePDF)

}
