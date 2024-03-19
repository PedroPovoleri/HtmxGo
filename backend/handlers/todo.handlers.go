package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"htmx/backend/services"

	"github.com/labstack/echo/v4"
)

/********** Handlers for Todo Views **********/

type TaskService interface {
	CreateTodo(t services.Todo) (services.Todo, error)
	GetAllTodos() ([]services.Todo, error)
	GetTodoById(t services.Todo) (services.Todo, error)
	UpdateTodo(t services.Todo) (services.Todo, error)
	DeleteTodo(t services.Todo) error
}

func NewTaskHandler(ts TaskService) *TaskHandler {
	return &TaskHandler{
		TodoServices: ts,
	}
}

type TaskHandler struct {
	TodoServices TaskService
}

func (th *TaskHandler) createTodoHandler(c echo.Context) error {
	isError = false

	if c.Request().Method == "POST" {
		todo := services.Todo{
			CreatedBy:   1,
			Title:       strings.Trim(c.FormValue("title"), " "),
			Description: strings.Trim(c.FormValue("description"), " "),
		}

		_, err := th.TodoServices.CreateTodo(todo)
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/todo")
	}
	return c.Redirect(http.StatusSeeOther, "/todo")

}

func (th *TaskHandler) todoListHandler(c echo.Context) error {
	isError = false

	todos, err := th.TodoServices.GetAllTodos()
	if err != nil {
		return err
	}

	return c.JSON(200, todos)
}

func (th *TaskHandler) updateTodoHandler(c echo.Context) error {

	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// t := services.Todo{
	// 	ID:        idParams,
	// 	CreatedBy: c.Get(1).(int),
	// }

	//todo, err := th.TodoServices.GetTodoById(t)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {

			return echo.NewHTTPError(
				echo.ErrNotFound.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}

	if c.Request().Method == "POST" {
		var status bool
		if c.FormValue("status") == "on" {
			status = true
		} else {
			status = false
		}

		todo := services.Todo{
			Title:       strings.Trim(c.FormValue("title"), " "),
			Description: strings.Trim(c.FormValue("description"), " "),
			Status:      status,
			ID:          idParams,
		}

		_, err := th.TodoServices.UpdateTodo(todo)
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/todo/list")
	}
	return c.Redirect(http.StatusSeeOther, "/todo/list")
}

func (th *TaskHandler) deleteTodoHandler(c echo.Context) error {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	t := services.Todo{
		CreatedBy: 1,
		ID:        idParams,
	}

	err = th.TodoServices.DeleteTodo(t)
	if err != nil {
		if strings.Contains(err.Error(), "an affected row was expected") {

			return echo.NewHTTPError(
				echo.ErrNotFound.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}

	return c.Redirect(http.StatusSeeOther, "/todo/list")
}
