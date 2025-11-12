package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/manish-npx/simple-go-echo/internal/config"
)

func ConnectDB(cfg *config.Config) *pgxpool.Pool {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Error Failed to connect DB %v", err)
	}

	error := pool.Ping(context.Background())
	if error != nil {
		log.Fatalf("Error! Can not ping %v", error)
	}

	fmt.Println("Database Connected to PostgreSQL Successfuly")

	return pool

}
