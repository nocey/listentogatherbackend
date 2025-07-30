package error_handler

import (
	"github.com/gofiber/fiber/v2"
	logger "github.com/listentogether/log"
)

func ErrorHandler (ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if fiberErr, ok := err.(*fiber.Error); ok {
		code = fiberErr.Code
		message = fiberErr.Message
	}

	logger.Warning(message)
	return ctx.SendStatus(code)
}