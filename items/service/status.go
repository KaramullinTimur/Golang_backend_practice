package service

import "backend/items/driver"

type StatusService struct {
	driver driver.Driver
}

func NewStatusService(driver driver.Driver) *StatusService {
	return &StatusService{driver}
}

func (s *StatusService) SetStatus() {
	// Some functions here
	s.driver.GetItems()
}
