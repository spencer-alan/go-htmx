package main

import (
	// "database/sql"
	// "fmt"
	"ghost/internal/common"
	"ghost/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	settings, err := common.ReadConfigToml("./ghost_config.toml")
	if err != nil {
		log.Fatal("Config failed to load: ", err)
	}

	// dbConnString := "user=postgres dbname=postgres password=password sslmode=disable"
	// db, err := sql.Open("postgres", dbConnString)
	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("Error opening database connection: %v\n", err)
	// }
	// defer db.Close()
	//
	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v\n", err)
	// } else {
	// 	fmt.Println("Successfully connected to the Database")
	// }

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.Init(e, settings)
	//routes.Init(e, settings, db)

	log.Println("Starting server on Port: 8000")
	err = e.Start(":8000")
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
