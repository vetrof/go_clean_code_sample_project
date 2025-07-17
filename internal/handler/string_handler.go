package handler

import (
	"go-fiber-app/internal/usecase/string"

	"github.com/gofiber/fiber/v2"
)

func NewStringHandler(uc *string.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result := uc.Do()
		return c.JSON(fiber.Map{"string": result})
	}
}
