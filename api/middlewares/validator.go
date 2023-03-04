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
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateBody[T any](dto T) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body T

		err := c.BodyParser(&body)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"errors": []string{"unable to parse request body"},
			})
		}

		var error_messages []string

		err = validate.Struct(body)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				error_messages = append(error_messages, fmt.Sprintf("%s %s %s", err.StructField(), err.Tag(), err.Param()))
			}

			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"errors": error_messages,
			})
		}

		c.Locals("dto", body)

		return c.Next()
	}
}

func ValidateQuery[T any](dto T) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var query T

		err := c.QueryParser(&query)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"errors": []string{"unable to parse request query"},
			})
		}

		var error_messages []string

		err = validate.Struct(query)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("ERRORS: ", err)
				
				error_messages = append(error_messages, fmt.Sprintf("%s %s %s", err.StructNamespace(), err.Tag(), err.Param()))
			}

			c.Status(http.StatusBadRequest)

			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"errors": error_messages,
			})
		}

		c.Locals("dto", query)
		
		return c.Next()
	}
}
