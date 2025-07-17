package handler

import (
	"go-fiber-app/internal/usecase/number"

	"github.com/gofiber/fiber/v2"
)

func NewNumberHandler(uc *number.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result := uc.Do()
		return c.JSON(fiber.Map{"number": result})
	}
}
