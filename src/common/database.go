package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InnitDB() {
	fmt.Printf("\nrun database from ENV...\n")

	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_URL")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s", userName, password, url)
	_, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("\nfound an error when try to connect to database:\n")
		log.Fatal(err)
		fmt.Print("\n")
	}

	fmt.Printf("\nSuccess run database from %s\n", url)
}
