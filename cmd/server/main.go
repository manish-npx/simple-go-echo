package main

import (
	"log"

	"github.com/manish-npx/simple-go-echo/internal/config"
	"github.com/manish-npx/simple-go-echo/internal/database"
	"github.com/manish-npx/simple-go-echo/internal/server"
)

func main() {
	log.Println("🚀 Starting application...")

	// Load configuration
	cfg := config.LoadConfig()

	// Setup database
	db := database.NewPostgres(cfg)
	defer db.Close()

	// Create and start server
	srv := server.NewServer(cfg, db)

	log.Println("🚀 Server running on:", cfg.Server.Addr)
	if err := srv.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
