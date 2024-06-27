package users

import (
	"context"
	"fmt"
	"get-on/database"
	"get-on/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateUser(c *fiber.Ctx) error {
	userData := new(model.User)
	if err := c.BodyParser(userData); err != nil {
		fmt.Println("error: ", err.Error())
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"error": "Invalid input data",
		})
	} else {
		createdUser, err := database.Client.User.
			Create().
			SetUsername(userData.Username).
			SetPassword(userData.Password).
			Save(context.TODO())

		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}

		c.SendStatus(201)
		return c.JSON(createdUser)
	}
}

func ReadUser(c *fiber.Ctx) error {
	client := database.Client
	u, err := client.User.Query().All(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(u)
}
