package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=admin dbname=dbuser sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}
