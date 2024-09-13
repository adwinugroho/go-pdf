package main

import (
	"go-pdf/route"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Go Generate PDF")

	app := fiber.New()

	route.InvoiceRoute(app, time.Duration(30)*time.Second)
	log.Fatal(app.Listen(":8000"))
}
