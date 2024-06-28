package users

import (
	"context"
	"fmt"
	"get-on/database"
	"get-on/ent"
	userField "get-on/ent/user"
	"get-on/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
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

func ReadAllUser(c *fiber.Ctx) error {
	client := database.Client
	user, err := client.User.Query().All(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(user)
}

func ReadUser(c *fiber.Ctx) error {
	userIdParam := c.Params("id")
	userId, err := strconv.Atoi(userIdParam)

	client := database.Client

	user, err := client.User.Query().
		Where(userField.IDEQ(userId)).
		Only(context.TODO())

	if err != nil {
		if ent.IsNotFound(err) {
			c.SendStatus(404)
			return c.JSON(fiber.Map{
				"error": "User not found",
			})
		}
		log.Fatal(err)
	}

	return c.JSON(user)
}
