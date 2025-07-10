package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewPostgresDb() *sql.DB {

	host := os.Getenv("DB_HOST") // Should be 'postgres'
	port := os.Getenv("DB_PORT") // Should be '5432'
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	log.Printf("DB env values -> host: %s, port: %s, user: %s", host, port, user)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
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
