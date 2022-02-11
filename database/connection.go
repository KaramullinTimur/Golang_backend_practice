package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDB() *sqlx.DB {
	// Connect to database
	str := "host=localhost port=5432 user=postgres password=00 dbname=aleyron sslmode=disable"
	db, err := sqlx.Open("postgres", str)
	if err != nil {
		log.Panicln("Fail to open database:", err)
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Panicln("Database connection failed:", err)
	}

	return db
}

func GetTestDB() *sqlx.DB {
	// Connect to database
	str := "host=localhost port=5432 user=postgres password=00 dbname=aleyron_test sslmode=disable"
	db, err := sqlx.Open("postgres", str)
	if err != nil {
		log.Panicln("Fail to open database:", err)
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Panicln("Database connection failed:", err)
	}

	return db
}
