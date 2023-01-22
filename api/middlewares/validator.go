package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag string `json:"tag"`
	Value string `json:"value"`
}

func ValidateBody[T any](dto T) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body T

		err := c.BodyParser(&body)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data": nil,
				"errors": []string{"unable to parse request body"},
			})
		}

		var error_messages []string

		err = validate.Struct(body)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				error_messages = append(error_messages, fmt.Sprintf("%s %s %s", err.StructNamespace(), err.Tag(), err.Param()))
			}

			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data": nil,
				"errors": error_messages,
			})
		}

		return c.Next()
	} 
}