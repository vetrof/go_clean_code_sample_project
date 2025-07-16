package main

import (
	"log"

	"random-service/internal/controller/http"
	"random-service/internal/repo/numberrepo"
	"random-service/internal/repo/stringrepo"
	numberUC "random-service/internal/usecase/number"
	stringUC "random-service/internal/usecase/string"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Wire dependencies
	stringRepo := stringrepo.NewMemoryRepo()
	numberRepo := numberrepo.NewMemoryRepo()

	stringService := stringUC.New(stringRepo)
	numberService := numberUC.New(numberRepo)

	http.RegisterRoutes(app, stringService, numberService)

	log.Fatal(app.Listen(":8080"))
}
