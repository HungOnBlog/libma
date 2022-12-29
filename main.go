package main

import (
	"os"

	"github.com/HungOnBlog/libma/infra/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
