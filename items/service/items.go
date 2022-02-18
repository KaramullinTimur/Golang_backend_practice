package service

import "backend/items/driver"

type ItemsService struct {
	driver driver.Driver
}

func NewItemsService(driver driver.Driver) *ItemsService {
	return &ItemsService{driver}
}

func (s *ItemsService) GetItems() {
	// Some functions here
	s.driver.GetItems()
}

func (s *ItemsService) AddItem() {

}
