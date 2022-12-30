package core

import (
	"github.com/gofiber/fiber/v2"

	libmaError "github.com/HungOnBlog/libma/core/errors"
)

func CustomLibmaErrorHandler(c *fiber.Ctx, err error) error {
	errorKey := err.Error()

	errorResponse := libmaError.GetLibmaError(errorKey)
	if errorResponse == (libmaError.LibmaError{}) {
		errorResponse = libmaError.GetLibmaError("InternalServerError")
	}

	return c.Status(errorResponse.Status).JSON(errorResponse)
}
