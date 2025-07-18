package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-fiber-app/internal/handler"
	numberRepo "go-fiber-app/internal/repository/number"
	stringRepo "go-fiber-app/internal/repository/string"
	"go-fiber-app/internal/router"
	numberUC "go-fiber-app/internal/usecase/number"
	stringUC "go-fiber-app/internal/usecase/string"
)

func main() {
	// init repositories
	numRepo := numberRepo.New()
	strRepo := stringRepo.New()

	// init usecases
	numberUseCase := numberUC.New(numRepo)
	stringUseCase := stringUC.New(strRepo)

	// init handlers
	numHandler := handler.NewNumberHandler(numberUseCase)
	strHandler := handler.NewStringHandler(stringUseCase)

	// create app and routes
	app := fiber.New()
	router.SetupRoutes(app, numHandler, strHandler)

	log.Fatal(app.Listen(":3000"))
}
