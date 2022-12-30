package core

import "github.com/gofiber/fiber/v2"

type IController interface {
	ApplyController(r fiber.Router)
}

func ApplyController(r fiber.Router, c IController) {
	c.ApplyController(r)
}
