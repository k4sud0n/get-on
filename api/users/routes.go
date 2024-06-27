package users

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Post("/users", CreateUser)
	route.Get("/users", ReadUser)
}
