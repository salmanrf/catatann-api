package common

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ExtractAuthorization(c *fiber.Ctx) string {
	authorization := c.Get("Authorization", "")

	if authorization == "" {
		return ""
	}

	auth_values := strings.Split(authorization, " ")

	return auth_values[len(auth_values) - 1]
}