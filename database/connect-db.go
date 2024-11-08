package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=inventory sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connStr)
	return db, err
}
