package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogStorage struct {
	DB *pgxpool.Pool
}

func NewBlogStorage(db *pgxpool.Pool) *BlogStorage {
	return &BlogStorage{DB: db}
}

// func (bs *BlogStorage) create(ctx context.Context) {

// }
