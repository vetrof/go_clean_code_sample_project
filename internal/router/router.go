package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, numberHandler fiber.Handler, stringHandler fiber.Handler) {
	// Простой health check эндпоинт
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	// Основные эндпоинты
	app.Get("/api/number", numberHandler)
	app.Get("/api/string", stringHandler)
}
