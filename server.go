package main

import (
	"context"
	"fmt"
	"log"

	"get-on/ent"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()

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
			CreateUser(client, user.Username, user.Password)
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

	// 회원조회
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		u, err := client.User.Query().All(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(u)
	})

	app.Listen(":3000")
}

func CreateUser(client *ent.Client, username string, password string) *ent.User {
	u, err := client.User.
		Create().
		SetUsername(username).
		SetPassword(password).
		Save(context.TODO())

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	return u
}
