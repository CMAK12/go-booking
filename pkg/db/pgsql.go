package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializePostgreSQL(connString *string) {
	var err error
	DB, err = sql.Open("postgres", *connString)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	err = createTables()
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}

	log.Println("database connection established")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		username VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(500) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		role VARCHAR(100) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS reservations (
		id UUID PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100),
		date DATE NOT NULL,
		guest_quantity INT NOT NULL,
		city VARCHAR(100) NOT NULL
	);`

	_, err := DB.Exec(query)
	return err
}
