package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/manish-npx/simple-go-echo/internal/model"
)

type TodoRepository struct {
	DB *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]model.Todo, error) {
	rows, err := r.DB.Query(ctx, "select * from todos order by id ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Todo])
	if err != nil {
		return nil, err
	}
	return todos, nil
}
