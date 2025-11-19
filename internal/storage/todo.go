package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/manish-npx/simple-go-echo/internal/models"
)

var ErrTodoNotFound = errors.New("todo not found")

type TodoStorage struct {
	DB *pgxpool.Pool
}

func NewTodoStorage(db *pgxpool.Pool) *TodoStorage {
	return &TodoStorage{DB: db}
}

func (s *TodoStorage) Create(ctx context.Context, todo *models.Todo) (int64, error) {
	var id int64
	err := s.DB.QueryRow(ctx,
		`INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id`,
		todo.Title, todo.Done,
	).Scan(&id)
	return id, err
}

func (s *TodoStorage) GetAll(ctx context.Context) ([]models.Todo, error) {
	rows, err := s.DB.Query(ctx, `SELECT id, title, done FROM todos ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	//find all the todos rows
	/*     todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Todo]) */

	return todos, nil
}

func (s *TodoStorage) GetByID(ctx context.Context, id int64) (*models.Todo, error) {
	var todo models.Todo
	err := s.DB.QueryRow(ctx,
		`SELECT id, title, done FROM todos WHERE id=$1`,
		id,
	).Scan(&todo.ID, &todo.Title, &todo.Done)

	if err != nil {
		return nil, ErrTodoNotFound
	}
	return &todo, nil
}

func (s *TodoStorage) Update(ctx context.Context, id int64, todo *models.Todo) (*models.Todo, error) {
	var updated models.Todo
	err := s.DB.QueryRow(ctx,
		`UPDATE todos SET title=$1, done=$2 WHERE id=$3 RETURNING id, title, done`,
		todo.Title, todo.Done, id,
	).Scan(&updated.ID, &updated.Title, &updated.Done)

	if err != nil {
		return nil, ErrTodoNotFound
	}
	return &updated, nil
}

func (s *TodoStorage) Delete(ctx context.Context, id int64) error {
	result, err := s.DB.Exec(ctx, `DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrTodoNotFound
	}
	return nil
}
