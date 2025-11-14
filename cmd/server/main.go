package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manish-npx/simple-go-echo/internal/config"
	"github.com/manish-npx/simple-go-echo/internal/db"
	"github.com/manish-npx/simple-go-echo/internal/handlers"
	"github.com/manish-npx/simple-go-echo/internal/repository"
	"github.com/manish-npx/simple-go-echo/internal/services"
)

func main() {

	log.Println("🚀 Main Function Started here ===>")
	//config done
	cfg := config.LoadConfig()

	//database connection
	pool := db.ConnectDB(cfg)

	defer pool.Close()

	//
	e := echo.New()

	//route
	//Setup Dependencies (Repository → Service → Handler)
	repo := repository.NewTodoRepository(pool)
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "Welcome to GO Echo"})
	})
	e.GET("/todos", handler.GetTodos)

	// create Echo web server
	log.Println("🚀 Server running on Add ===>", cfg.Server.Addr)

	//server
	err := e.Start(cfg.Server.Addr)
	e.Logger.Fatal(err) // start server on given port

}
