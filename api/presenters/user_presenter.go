package presenters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/pkg/models"
)

func UserSuccessResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data": data,
		"errors": []string{},
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data": nil,
		"errors": []string{err.Error()},
	}
}

func UserCustomErrorResponse(cust_err *models.CustomHttpErrors) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data": nil,
		"errors": cust_err.Messages,
	}
}