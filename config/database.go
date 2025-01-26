package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() *sql.DB {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// input from .env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connected to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database")

	return DB

}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
