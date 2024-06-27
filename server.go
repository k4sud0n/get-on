package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 회원가입
	app.Post("/api/v1/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			fmt.Println("error: ", err.Error())
			c.SendStatus(400)
			return c.JSON(fiber.Map{
				"error": "Invalid input data",
			})
		} else {
			c.SendStatus(201)
			return c.JSON(user)
		}
	})

	// 로그인
	app.Post("/api/v1/auth/login", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			fmt.Println("error: ", err.Error())
			c.SendStatus(401)
			return c.JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		} else {
			c.SendStatus(200)
			return c.JSON(fiber.Map{
				"token": "jwt_token",
			})
		}
	})

	app.Listen(":3000")
}
