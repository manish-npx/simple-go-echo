package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OK(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

func Created(c echo.Context, data any) error {
	return c.JSON(http.StatusCreated, data)
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func BadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": msg})
}

func NotFound(c echo.Context, msg string) error {
	return c.JSON(http.StatusNotFound, map[string]string{"error": msg})
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}

func CustomErrorHandler(err error, c echo.Context) {
	// Check if it's an echo HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, map[string]any{
			"error": he.Message,
		})
		return
	}

	// Default to internal server error
	c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}
