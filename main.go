package main

import (
	"os"

	"github.com/HungOnBlog/libma/app/borrowers"
	"github.com/HungOnBlog/libma/core"
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
	core.AutoMigrate(
		borrowers.GetBorrowerRepo(),
	)

	app := fiber.New(fiber.Config{
		Prefork:       os.Getenv("PREFORK") == "true",
		CaseSensitive: true,
		AppName:       "Libma",
		ErrorHandler:  core.CustomLibmaErrorHandler,
	})

	// Middleware
	core.ApplyMiddleware(app)

	apiV1 := app.Group("v1")

	// Controller
	core.ApplyController(apiV1, &borrowers.BorrowerController{})

	// Swagger
	swagger.ApplySwaggerController(apiV1)

	app.Listen(":" + os.Getenv("PORT"))
}
