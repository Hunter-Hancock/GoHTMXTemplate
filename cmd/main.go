package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = NewTemplate()

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
	InitDB(config)

	e.GET("/", HandleHome)
	e.GET("/orders", HandleGetAllOrders)
	e.GET("/createOrder", HandleCreateOrder)

	e.POST("/createOrder", HandleCreateOrder)
	e.DELETE("/deleteOrder/:id", HandleDeleteOrder)

	e.Logger.Fatal(e.Start(":3000"))
}
