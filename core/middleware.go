package core

import (
	"time"

	"github.com/HungOnBlog/libma/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func ApplyMiddleware(a *fiber.App) {
	applyRecover(a)
	applyRequestId(a)
	applyCache(a)
}

func applyRecover(a *fiber.App) {
	a.Use(recover.New())
}

func applyRequestId(a *fiber.App) {
	a.Use(requestid.New(requestid.Config{
		Header: "requestId",
		Generator: func() string {
			return utils.GenRequestId()
		},
	}))
}

func applyCache(a *fiber.App) {
	a.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
}
