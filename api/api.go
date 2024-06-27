package api

import (
	"get-on/api/auth"
	"get-on/api/users"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	auth.Routes(v1)
	users.Routes(v1)
}
