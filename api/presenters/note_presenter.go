package presenters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/pkg/models"
)

func NoteSuccessResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func NoteErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}

func NoteCustomErrorResponse(cust_err *models.CustomHttpErrors) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data": nil,
		"errors": cust_err.Messages,
	}
}