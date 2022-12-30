package main

import (
	"os"

	"github.com/HungOnBlog/libma/infra/db"
	"github.com/HungOnBlog/libma/infra/swagger"
	"github.com/gofiber/fiber/v2"
)

// @title Libma API Specification
// @version 1.0
// @description This is API specification for Libma
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	db.InitDb()
	app := fiber.New()

	apiV1 := app.Group("v1")

	swagger.ApplySwaggerController(apiV1)

	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
