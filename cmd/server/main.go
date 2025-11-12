package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/manish-npx/simple-go-echo/internal/config"
	"github.com/manish-npx/simple-go-echo/internal/db"
)

func main() {

	log.Println("🚀 Main Function Started ===>")
	//config done
	cfg := config.LoadConfig()

	//database connection
	pool := db.ConnectDB(cfg)

	defer pool.Close()

	//route

	// create Echo web server

	e := echo.New()
	log.Println("🚀 Server running on Add ===>", cfg.Server.Addr)
	//server
	err := e.Start(cfg.Server.Addr)
	e.Logger.Fatal(err) // start server on given port

}
