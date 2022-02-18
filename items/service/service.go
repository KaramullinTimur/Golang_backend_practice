package service

import (
	"backend/items/driver"
)

type Items interface {
	GetItems()
	AddItem()
}

type Status interface {
	SetStatus()
}

type Service struct {
	Items
	Status
}

func NewService(driver driver.Driver) *Service {
	return &Service{
		Items:  NewItemsService(driver),
		Status: NewStatusService(driver),
	}
}
