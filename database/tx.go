package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Transaction(ctx context.Context, db *sqlx.DB, fn func(tx sqlx.ExtContext) error) (err error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	defer func() {
		if r := recover(); r != nil || err != nil {
			_ = tx.Rollback()
			if r != nil {
				err = fmt.Errorf("transaction panic: %v", r)
			} else {
				err = fmt.Errorf("transaction erred: %w", err)
			}
			return
		}
		err = tx.Commit()
		if err != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("transaction fail to commit: %w", err)
		}
	}()

	return fn(tx)
}
