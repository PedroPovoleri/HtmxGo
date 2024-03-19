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

	ts = services.TodoServices(services.Todo{}, store)
	th = handlers.NewTaskHandler(ts)

	handlers.SetupRoutes(e, th)

	e.Logger.Fatal(e.Start(":1323"))
}
