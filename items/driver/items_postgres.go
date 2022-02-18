package driver

import "github.com/jmoiron/sqlx"

type itemsDriver struct {
	db sqlx.DB
}

func NewPostgresItemsDriver(db sqlx.DB) Items {
	return &itemsDriver{db}
}

func (d *itemsDriver) GetItems() {
	d.db.Query()
}

func (d *itemsDriver) AddItem() {
	d.db.Query()
}
