package api

import (
	"backend/items/service"

	"github.com/gofiber/fiber"
)

type api struct {
	service service.Service
}

func NewHttpApi(service service.Service) API {
	return &api{service}
}

func (a *api) HandleHttpRequest(r fiber.Router) {
	r.Get("/items", a.getItems)
	r.Post("/items/:id", a.addItem)
	r.Post("/status/:id", a.setStatus)
}
