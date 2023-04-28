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

type ActivityHandler struct {
	AcitivityService service.ActivityService
}

func NewActivityHandler(activityService *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{AcitivityService: *activityService}
}

func (handler *ActivityHandler) FindAll(c *fiber.Ctx) error {
	activities, err := handler.AcitivityService.FindAll(c.Context())
	if err != nil {
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", activities)
}

func (handler *ActivityHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	activity, err := handler.AcitivityService.FindByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Activity with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", activity)
}

func (handler *ActivityHandler) Create(c *fiber.Ctx) error {
	var activityRequest request.ActivityCreateRequest
	if err := c.BodyParser(&activityRequest); err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", err.Error())
	}

	if activityRequest.Title == "" {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", "title cannot be null")
	}

	activity, err := handler.AcitivityService.Create(c.Context(), &model.Activity{
		Title: activityRequest.Title,
		Email: activityRequest.Email,
	})
	if err != nil {
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusCreated, "Success", "Success", activity)
}

func (handler *ActivityHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	var activityRequest request.ActivityUpdateRequest
	if err := c.BodyParser(&activityRequest); err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	if activityRequest.Title == "" {
		return response.ReturnError(c, http.StatusBadRequest, "Bad Request", "title cannot be null")
	}

	activity, err := handler.AcitivityService.Update(c.Context(), &model.Activity{
		ActivityID: id,
		Title:      activityRequest.Title,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Activity with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", activity)
}

func (handler *ActivityHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ReturnError(c, http.StatusBadRequest, "Error", err.Error())
	}

	err = handler.AcitivityService.Delete(c.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Activity with ID %d Not Found", id)
			return response.ReturnError(c, http.StatusNotFound, "Not Found", errorString)
		}
		return response.ReturnError(c, http.StatusInternalServerError, "Error", err.Error())
	}

	return response.ReturnSuccess(c, http.StatusOK, "Success", "Success", nil)
}
