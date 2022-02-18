package driver

import "github.com/jmoiron/sqlx"

type statusDriver struct {
	db sqlx.DB
}

func NewPostgresStatusDriver(db sqlx.DB) Status {
	return &statusDriver{db}
}

func (d *statusDriver) SetStatus() {
	d.db.Query()
}
