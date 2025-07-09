package db

import (
	"database/sql"
	"fmt"
	"os"
)

func NewPostgresDb() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("cannot ping postgres: %v", err))
	}

	return db
}
