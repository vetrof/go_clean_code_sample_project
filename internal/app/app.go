package app

import (
	v1 "random-service/internal/controller/v1"
	numberrepo "random-service/internal/repo/number"
	stringrepo "random-service/internal/repo/string"
	"random-service/internal/usecase"
	"random-service/pkg/httpserver"
	"random-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	log *logger.Logger
}

func NewApp(log *logger.Logger) *App {
	return &App{log: log}
}

func (a *App) Run() error {
	app := fiber.New()

	// Repositories
	stringRepo := stringrepo.NewMemoryRepo()
	numberRepo := numberrepo.NewMemoryRepo()

	// Usecases
	stringUC := usecase.NewStringUseCase(stringRepo)
	numberUC := usecase.NewNumberUseCase(numberRepo)

	// Controllers
	v1.RegisterRoutes(app, stringUC, numberUC)

	// HTTP Server
	return httpserver.New(app).Start()
}
