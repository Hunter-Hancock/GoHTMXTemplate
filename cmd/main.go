package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DBHOST")
	dbtype := os.Getenv("DBTYPE")
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")

	config := NewDBConfig(host, dbtype, dbname, dbuser, dbpass)
	store, err := InitDB(config)
	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	server := NewServer(store)
	server.Run()
}
