package items

import (
	"backend/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type HttpService interface {
	GetItems(c *fiber.Ctx) error
}

type httpService struct {
	db *sqlx.DB
}

func NewHttpService(db *sqlx.DB) HttpService {
	return &httpService{db}
}

func (h *httpService) GetItems(c *fiber.Ctx) error {
	ctx := c.Context()
	err := database.Transaction(ctx, h.db, func(tx sqlx.ExtContext) error {
		err := GetItemsFromDB(h.db)
		if err != nil {
			fmt.Println(fmt.Errorf("ERROR: failed to get items from database: %w", err))
			return err
		}
		return c.SendString("Hello!!")
	})
	if err != nil {
		return fmt.Errorf("ERROR: transaction failed: %w", err)
	}
	return nil

}

// Structs

type Item struct {
	ID     int
	Name   string
	Age    int
	Active bool
}
