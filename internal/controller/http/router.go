package http

import (
	numberUC "random-service/internal/usecase/number"
	stringUC "random-service/internal/usecase/string"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, stringUsecase *stringUC.UseCase, numberUsecase *numberUC.UseCase) {
	api := app.Group("/api/v1/random")

	api.Get("/string", func(c *fiber.Ctx) error {
		res, err := stringUsecase.Get()
		if err != nil {
			return c.Status(500).SendString("internal error")
		}
		return c.JSON(res)
	})

	api.Get("/number", func(c *fiber.Ctx) error {
		res, err := numberUsecase.Get()
		if err != nil {
			return c.Status(500).SendString("internal error")
		}
		return c.JSON(res)
	})
}
