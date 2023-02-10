package presenters

import (
	"github.com/gofiber/fiber/v2"
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
