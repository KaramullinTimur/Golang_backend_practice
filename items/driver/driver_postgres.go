package driver

import "github.com/jmoiron/sqlx"

type driver struct {
	Items
	Status
}

func NewPostgresDriver(db *sqlx.DB) Driver {
	return &driver{
		Items:  NewPostgresItemsDriver(db),
		Status: NewPostgresStatusDriver(db),
	}
}
