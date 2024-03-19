package main

import (
	"htmx/backend/db"
	"htmx/backend/handlers"
	"htmx/backend/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "frontend")

	handlers.SetupLogs(e)

	store, err := db.NewStore("PoC")
	if err != nil {
		e.Logger.Fatalf("failed to create store: %s", err)
	}

	var todoService = services.TodoServices{
		Todo:      services.Todo{},
		TodoStore: store,
	}

	var taskHandler = handlers.NewTaskHandler(todoService)

	handlers.SetupRoutes(e, taskHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
