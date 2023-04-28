package internal

import (
	"database/sql"
	"github.com/arvians-id/go-todo-list/internal/handler"
	"github.com/arvians-id/go-todo-list/internal/repository"
	"github.com/arvians-id/go-todo-list/internal/service"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(c fiber.Router, db *sql.DB) {
	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(activityRepository)
	activityHandler := handler.NewActivityHandler(activityService)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	// Activity Routes
	c.Get("/activity-groups", activityHandler.FindAll)
	c.Get("/activity-groups/:id", activityHandler.FindByID)
	c.Post("/activity-groups", activityHandler.Create)
	c.Patch("/activity-groups/:id", activityHandler.Update)
	c.Delete("/activity-groups/:id", activityHandler.Delete)

	// Todo Routes
	c.Get("/todo-items", todoHandler.FindAll)
	c.Get("/todo-items/:id", todoHandler.FindByID)
	c.Post("/todo-items", todoHandler.Create)
	c.Patch("/todo-items/:id", todoHandler.Update)
	c.Delete("/todo-items/:id", todoHandler.Delete)
}
