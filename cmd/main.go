package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var db sqlx.DB

func main() {
	// aaa
	// s := aaa.ApiFunc()
	// fmt.Println(s)

	// Connect to database
	// db := database.GetDB()

	// Initialize fiber app
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	// driver := driver.NewPostgresDriver(121)
	// service := service.NewService(driver)
	// api := api.NewHttpApi(service)

	// api.HandleHttpRequest()

	// Old items
	// srvc := items.NewHttpService(db)
	// app.Get("/", srvc.GetItems)

	err := app.Listen(":3001")
	if err != nil {
		log.Panicln("Failed to listen 3001:", err)
	}
}

func Max(num int) bool {
	if num > 0 {
		return true
	}
	return false
}
