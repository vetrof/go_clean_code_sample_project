package v1

import (
	"random-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App, sUC usecase.StringUseCase, nUC usecase.NumberUseCase) {
	v1 := router.Group("/api/v1")

	v1.Get("/random/string", func(c *fiber.Ctx) error {
		res, err := sUC.Get()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	v1.Get("/random/number", func(c *fiber.Ctx) error {
		res, err := nUC.Get()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})
}
