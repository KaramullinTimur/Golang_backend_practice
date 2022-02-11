package main

import (
	"backend/aaa"
	"backend/database"
	"backend/items"
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var db sqlx.DB

func main() {
	s := aaa.ApiFunc()
	fmt.Println(s)

	// Connect to database
	db := database.GetDB()

	// Initialize fiber app
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	srvc := items.NewHttpService(db)
	app.Get("/", srvc.GetItems)

	err := app.Listen(":3000")
	if err != nil {
		log.Panicln("Failed to listen 3000:", err)
	}
}
