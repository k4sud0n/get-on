package users

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Post("/users", CreateUser)
	route.Get("/users", ReadAllUser)
	route.Get("/users/:id", ReadUser)
}
