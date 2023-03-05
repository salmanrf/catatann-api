package middlewares

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/presenters"
	"github.com/salmanfr/catatann-api/pkg/common"
	"github.com/salmanfr/catatann-api/pkg/models"
)

func AuthorizationGuard(c *fiber.Ctx) error {
	access_token := common.ExtractAuthorization(c)

	if access_token == "" {
		c.Status(http.StatusUnauthorized)
		
		return c.JSON(presenters.UserCustomErrorResponse(models.CreateCustomHttpError(http.StatusUnauthorized, "unauthorized")))
	}

	claims, err := common.VerifyJwt(access_token, os.Getenv("USER_ACCESS_TOKEN_JWT_SECRET"))

	if err != nil {
		c.Status(http.StatusUnauthorized)
		
		return c.JSON(presenters.UserCustomErrorResponse(models.CreateCustomHttpError(http.StatusUnauthorized, err)))
	}

	c.Locals("decoded", claims)

	return c.Next()
}