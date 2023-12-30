package error_handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	c.Status(code).JSON(fiber.Map{
		"isOk": false,
		"data": fiber.Map{
			"message": e.Error(),
		},
	})
	return nil
}
