package swagger

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "github.com/HungOnBlog/libma/docs"
)

func ApplySwaggerController(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault)
}
