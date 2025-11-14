package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manish-npx/simple-go-echo/internal/services"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (t *TodoHandler) GetTodos(c echo.Context) error {
	todos, err := t.service.ListTodo(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}
