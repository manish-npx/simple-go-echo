package services

import (
	"context"

	"github.com/manish-npx/simple-go-echo/internal/model"
	"github.com/manish-npx/simple-go-echo/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (t *TodoService) ListTodo(ctx context.Context) ([]model.Todo, error) {
	return t.repo.GetAll(ctx)

}
