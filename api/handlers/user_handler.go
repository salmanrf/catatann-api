package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/presenters"
	"github.com/salmanfr/catatann-api/pkg/common"
	"github.com/salmanfr/catatann-api/pkg/models"
	"github.com/salmanfr/catatann-api/pkg/user"
)

func GetSelf(s user.Service) fiber.Handler {
	return func (c *fiber.Ctx) error {
		access_token := common.ExtractAuthorization(c)

		if access_token == "" {
			c.Status(http.StatusUnauthorized)
			
			return c.JSON(presenters.UserCustomErrorResponse(models.CreateCustomHttpError(http.StatusUnauthorized, "unauthorized")))
		}

		res, err := s.GetSelf(access_token)

		if err != nil {
			c.Status(err.Code)
			
			return c.JSON(presenters.UserCustomErrorResponse(err))
		}

		return c.JSON(presenters.UserSuccessResponse(res))
	}
}

func GoogleSignin(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		path_url := "/"

		if state := c.Query("state"); state != "" {
			path_url = state
		}

		res, err := s.GoogleSignin(code)

		if err != nil {
			c.Status(err.Code)
			
			return c.JSON(presenters.UserCustomErrorResponse(err))
		}

		c.Cookie(&fiber.Cookie{
			Name: "ctnn_refresh_token",
			Value: res.RefreshToken,
			Expires: time.Now().Add(time.Hour * 24 * 365),
			Path: "/",
			Domain: "",
			Secure: true,
			HTTPOnly: true,
		})

		return c.Redirect(fmt.Sprintf("http://localhost:5173%s", path_url))
	}
}

func Signup(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto, valid := c.Locals("dto").(models.SignupDto)

		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.UserErrorResponse(errors.New("unable to parse request body")))
		}

		result, err := s.Signup(dto)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.UserErrorResponse(err))
		}

		c.Status(http.StatusOK)
		
		return c.JSON(result)
	}
}

func Signin(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto := c.Locals("dto").(models.SigninDto)

		res, err := s.Signin(dto)

		if err != nil {			
			c.Status(err.Code)
			
			return c.JSON(presenters.UserCustomErrorResponse(err))
		}

		c.Cookie(&fiber.Cookie{
			Name: "ctnn_refresh_token",
			Value: res.RefreshToken,
			Expires: time.Now().Add(time.Hour * 24 * 365),
			Path: "/",
			Domain: "",
			Secure: true,
			HTTPOnly: true,
		})
		
		res.RefreshToken = ""

		return c.JSON(presenters.UserSuccessResponse(res))
	}
}

func Signout(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refresh_token := c.Cookies("ctnn_refresh_token", "")
		
		fmt.Println("ctnn_refresh_token", refresh_token)
		
		claims, err := common.VerifyJwt(refresh_token, os.Getenv("USER_REFRESH_TOKEN_JWT_SECRET"))

		fmt.Println("claims", claims)
		fmt.Println("err", err)
		
		if err == nil {
			fmt.Println("sub", claims["sub"])
			
			_ = s.InvalidateRefreshToken(claims["sub"].(string))
		}
		
		c.Cookie(&fiber.Cookie{
			Name: "ctnn_refresh_token",
			Value: "",
			Expires: time.Now().Add(time.Hour * -24),
			Path: "/",
			Domain: "",
			Secure: true,
			HTTPOnly: true,
		})
		
		return c.JSON(presenters.UserSuccessResponse(nil))
	}
}

func GetRefreshToken(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refresh_token := c.Cookies("ctnn_refresh_token", "")

		if refresh_token == "" {
			c.Status(http.StatusUnauthorized)
			
			return c.JSON(presenters.UserCustomErrorResponse(models.CreateCustomHttpError(http.StatusUnauthorized, "session has expired")))
		}

		res, err := s.RefreshToken(refresh_token)

		if err != nil {
			c.Status(err.Code)
			
			return c.JSON(presenters.UserCustomErrorResponse(err))
		}

		c.Cookie(&fiber.Cookie{
			Name: "ctnn_refresh_token",
			Value: res.RefreshToken,
			Expires: time.Now().Add(time.Hour * 24 * 365),
			Path: "/",
			Domain: "",
			Secure: true,
			HTTPOnly: true,
		})
		
		res.RefreshToken = ""

		return c.JSON(presenters.UserSuccessResponse(res))
	}
}