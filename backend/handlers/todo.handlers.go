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

type TodoServices interface {
	CreateTodo(t services.Todo) (services.Todo, error)
	GetAllTodos() ([]services.Todo, error)
	GetTodoById(t services.Todo) (services.Todo, error)
	UpdateTodo(t services.Todo) (services.Todo, error)
	DeleteTodo(t services.Todo) error
}

func NewTaskHandler(ts TodoServices) *TaskHandler {
	return &TaskHandler{
		TodoServices: ts,
	}
}

type TaskHandler struct {
	TodoServices TodoServices
}

func (th *TaskHandler) GetAllTodos(c echo.Context) error {
	todos, _ := th.TodoServices.GetAllTodos()
	var html = ""
	for i := 0; i <= len(todos)-1; i++ {
		html += "<li class=' list-group-item d-flex justify-content-between align-items-center' ><span>" + todos[i].Title + "</span><i class='far fa-trash-alt delete'></i>"
	}

	if len(todos) == 0 {
		return c.HTML(http.StatusNotFound, "")
	}

	return c.HTML(http.StatusOK, html)
}

func (th *TaskHandler) createTodoHandler(c echo.Context) error {
	todo := services.Todo{
		CreatedBy:   1,
		Title:       strings.Trim(c.FormValue("title"), " "),
		Description: strings.Trim(c.FormValue("description"), " "),
	}

	_, err := th.TodoServices.CreateTodo(todo)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (th *TaskHandler) todoListHandler(c echo.Context) error {
	todos, err := th.TodoServices.GetAllTodos()
	if err != nil {
		return err
	}

	return c.JSON(200, todos)
}

func (th *TaskHandler) updateTodoHandler(c echo.Context) error {

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

	_, err = th.TodoServices.UpdateTodo(todo)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
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

	return c.Redirect(http.StatusSeeOther, "/")
}
