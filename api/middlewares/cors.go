package middlewares

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Cors(c * fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", os.Getenv("CATATANN_CLIENT_URL"))
	c.Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	c.Set("Access-Control-Allow-Credentials", "true")

	if c.Method() == "OPTIONS" {
		c.Status(http.StatusOK)
		
		return c.Send([]byte{})
	} 
	
	return c.Next()
}