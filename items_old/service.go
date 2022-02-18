package items_old

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetItemsFromDB(db *sqlx.DB) error {
	var item string

	err := db.Get(&item, "SELECT name FROM atest WHERE id=3")
	if err != nil {
		return fmt.Errorf("failed SQL query: %w", err)
	}

	fmt.Println(item)
	return nil
}
