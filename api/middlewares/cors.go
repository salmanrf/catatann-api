package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Cors(c * fiber.Ctx) error {
	// fmt.Println("cookies", c.Cookies("ctnn_access_token"))
	// fmt.Println("origin", c.GetReqHeaders()["Origin"])
	// fmt.Println("METHOD", c.Method())
	
	c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	c.Set("Access-Control-Allow-Credentials", "true")

	if c.Method() == "OPTIONS" {
		c.Status(http.StatusOK)
		
		return c.Send([]byte{})
	} 
	
	return c.Next()
}