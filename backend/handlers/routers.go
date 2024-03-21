package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo, th *TaskHandler) {

	e.GET("/TODO", func(c echo.Context) error {
		return th.GetAllTodos(c)
	})

	e.POST("/TODO", func(c echo.Context) (err error) {
		return th.createTodoHandler(c)
	})

	e.DELETE("/TODO", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
