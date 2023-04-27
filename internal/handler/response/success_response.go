package response

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
)

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NullResponse struct {
}

func ReturnSuccess(c *fiber.Ctx, statusCode int, status string, message string, data interface{}) error {
	var null NullResponse
	if data == nil {
		data = null
	}
	if reflect.ValueOf(data).Kind() == reflect.Slice && reflect.ValueOf(data).IsNil() {
		data = []string{}
	}
	return c.Status(statusCode).JSON(SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
