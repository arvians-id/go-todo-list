package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/arvians-id/go-todo-list/internal/handler/request"
	"github.com/arvians-id/go-todo-list/internal/handler/response"
	"github.com/arvians-id/go-todo-list/internal/model"
	"github.com/arvians-id/go-todo-list/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TodoHandler struct {
	TodoService service.TodoServiceContract
}

func NewTodoHandler(todoService service.TodoServiceContract) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

func (handler *TodoHandler) FindAll(c *fiber.Ctx) error {
	queryGroupID := c.QueryInt("activity_group_id")
	todos, err := handler.TodoService.FindAll(c.Context(), queryGroupID)
	if err != nil {
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", todos)
}

func (handler *TodoHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	todo, err := handler.TodoService.FindByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Todo with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", todo)
}

func (handler *TodoHandler) Create(c *fiber.Ctx) error {
	var todoRequest request.CreateTodoRequest
	if err := c.BodyParser(&todoRequest); err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", err.Error())
	}

	if todoRequest.Title == "" {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", "title cannot be null")
	}

	if todoRequest.Priority == "" {
		todoRequest.Priority = "very-high"
	}

	if todoRequest.ActivityGroupID == 0 {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", "activity_group_id cannot be null")
	}

	todo, err := handler.TodoService.Create(c.Context(), &model.Todo{
		Title:           todoRequest.Title,
		ActivityGroupID: todoRequest.ActivityGroupID,
		IsActive:        true,
		Priority:        todoRequest.Priority,
	})
	if err != nil {
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusCreated, "Success", "Success", todo)
}

func (handler *TodoHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	var todoRequest request.UpdateTodoRequest
	if err := c.BodyParser(&todoRequest); err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", err.Error())
	}

	todo, err := handler.TodoService.Update(c.Context(), &model.Todo{
		TodoID:   id,
		Title:    todoRequest.Title,
		IsActive: todoRequest.IsActive,
		Priority: todoRequest.Priority,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Todo with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", todo)
}

func (handler *TodoHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	err = handler.TodoService.Delete(c.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Todo with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", nil)
}
