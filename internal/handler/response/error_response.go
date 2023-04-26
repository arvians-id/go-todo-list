package response

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ReturnError(c *fiber.Ctx, statusCode int, status string, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Status:  status,
		Message: message,
	})
}
