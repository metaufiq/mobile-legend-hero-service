package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InnitDB() *sql.DB {
	if db != nil {
		fmt.Printf("\nuse existing connection..\n")
		return db
	}
	fmt.Printf("\nrun database from ENV...\n")

	dbLocation := os.Getenv("DB_LOCATION")

	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", dbLocation, dbPort,
		userName, password, dbName)
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("\nfound an error when try to connect to database:\n")
		log.Fatal(err)
		fmt.Print("\n")
	}

	db = conn

	return db
}
