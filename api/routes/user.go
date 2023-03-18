package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/handlers"
	"github.com/salmanfr/catatann-api/api/middlewares"
	"github.com/salmanfr/catatann-api/pkg/models"
	"github.com/salmanfr/catatann-api/pkg/user"
)

func UserRoutes(app fiber.Router, s user.Service) {
	app.Get("/google-oauth/callback", handlers.GoogleSignin(s))
	app.Post("/signup", middlewares.ValidateBody(models.SignupDto{}), handlers.Signup(s))
	app.Post("/signin", middlewares.ValidateBody(models.SigninDto{}), handlers.Signin(s))
	app.Get("/extension-refresh", handlers.GetExtensionRefreshToken(s))
	app.Get("/signout", handlers.Signout(s))
	app.Get("/self", handlers.GetSelf(s))
	app.Get("/refresh", handlers.GetRefreshToken(s))
}