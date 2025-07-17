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
	// Инициализация зависимостей
	numRepo := numberRepo.New()
	strRepo := stringRepo.New()

	numUseCase := numberUC.New(numRepo)
	strUseCase := stringUC.New(strRepo)

	numHandler := handler.NewNumberHandler(numUseCase)
	strHandler := handler.NewStringHandler(strUseCase)

	// Создание Fiber-приложения и маршрутов
	app := fiber.New()
	router.SetupRoutes(app, numHandler, strHandler)

	// Запуск сервера
	log.Fatal(app.Listen(":3000"))
}
