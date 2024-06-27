package auth

import (
	"fmt"
	"get-on/model"
	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	user := new(model.User)
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
}
