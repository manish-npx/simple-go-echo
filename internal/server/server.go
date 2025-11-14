package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manish-npx/simple-go-echo/internal/config"
	"github.com/manish-npx/simple-go-echo/internal/http/handlers"
	"github.com/manish-npx/simple-go-echo/internal/storage"
	"github.com/manish-npx/simple-go-echo/internal/utils/response"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewServer(cfg *config.Config, db *pgxpool.Pool) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = response.CustomErrorHandler

	// Initialize storage and handlers
	todoStorage := storage.NewTodoStorage(db)
	todoHandler := handlers.NewTodoHandler(todoStorage)

	// Routes
	api := e.Group("/api")
	api.GET("/todos", todoHandler.GetAll)
	api.POST("/todos", todoHandler.Create)
	api.GET("/todos/:id", todoHandler.GetByID)
	api.PUT("/todos/:id", todoHandler.Update)
	api.DELETE("/todos/:id", todoHandler.Delete)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(s.cfg.Server.Addr)
}
