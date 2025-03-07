package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {

	// Fetch database details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct the data source name (DSN)
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Open the connection to the database
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Create the doctors table if it does not exist
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS doctors (
        id VARCHAR(36) PRIMARY KEY,
        first_name VARCHAR(100),
        last_name VARCHAR(100),
        specialty VARCHAR(100),
        phone_number VARCHAR(15),
        email VARCHAR(100)
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	log.Println("Database initialized and doctors table created successfully.")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
