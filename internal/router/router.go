package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, numberHandler fiber.Handler, stringHandler fiber.Handler) {
	app.Get("/api/number", numberHandler)
	app.Get("/api/string", stringHandler)
}
