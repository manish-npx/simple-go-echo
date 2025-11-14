package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manish-npx/simple-go-echo/internal/models"
	"github.com/manish-npx/simple-go-echo/internal/storage"
	"github.com/manish-npx/simple-go-echo/internal/utils/response"
)

type TodoHandler struct {
	storage *storage.TodoStorage
}

func NewTodoHandler(storage *storage.TodoStorage) *TodoHandler {
	return &TodoHandler{storage: storage}
}

func (h *TodoHandler) GetAll(c echo.Context) error {
	todos, err := h.storage.GetAll(c.Request().Context())
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, todos)
}

func (h *TodoHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	todo, err := h.storage.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.NotFound(c, "Todo not found")
	}
	return response.OK(c, todo)
}

func (h *TodoHandler) Create(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if todo.Title == "" {
		return response.BadRequest(c, "Title is required")
	}

	id, err := h.storage.Create(c.Request().Context(), &todo)
	if err != nil {
		return response.InternalServerError(c, err)
	}

	todo.ID = id
	return response.Created(c, todo)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if todo.Title == "" {
		return response.BadRequest(c, "Title is required")
	}

	updated, err := h.storage.Update(c.Request().Context(), id, &todo)
	if err != nil {
		return response.NotFound(c, "Todo not found")
	}

	return response.OK(c, updated)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	err = h.storage.Delete(c.Request().Context(), id)
	if err != nil {
		return response.NotFound(c, "Todo not found")
	}
	return response.NoContent(c)
}
