package response

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnSuccess(c *fiber.Ctx, statusCode int, status string, message string, data interface{}) error {
	return c.Status(statusCode).JSON(SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
