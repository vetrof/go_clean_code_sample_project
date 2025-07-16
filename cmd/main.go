package main

import (
	"log"
	"place_service/internal/place/handler"
	"place_service/internal/place/repository"
	"place_service/internal/place/router"
	"place_service/internal/place/service"
)

func main() {
	// Инициализация логгера
	log.Println("Starting Place Service...")

	// === СЛОЙ РЕПОЗИТОРИЕВ (DATA LAYER) ===
	// Создаем репозитории для работы с данными
	nearbyRepo := repository.MokNearbyRepository()
	popularRepo := repository.NewPopularRepository()
	placeRepo := repository.NewPlaceRepository()
	log.Println("✓ Repositories initialized")

	// === СЛОЙ СЕРВИСОВ (BUSINESS LOGIC LAYER) ===
	// Создаем сервисы, которые содержат бизнес-логику
	// Сервисы зависят от интерфейсов репозиториев, а не от конкретных реализаций
	nearbyService := service.NewNearbyService(nearbyRepo)
	popularService := service.NewPopularService(popularRepo)
	placeService := service.NewPlaceService(placeRepo)
	log.Println("✓ Services initialized")

	// === СЛОЙ ХЕНДЛЕРОВ (PRESENTATION LAYER) ===
	// Создаем хендлеры для обработки HTTP запросов
	// Хендлеры зависят от интерфейсов сервисов, а не от конкретных реализаций
	nearbyHandler := handler.NewNearbyHandler(nearbyService)
	popularHandler := handler.NewPopularHandler(popularService)
	placeHandler := handler.NewPlaceHandler(placeService)
	log.Println("✓ Handlers initialized")

	// === СЛОЙ РОУТЕРА (ROUTING LAYER) ===
	// Создаем роутер для настройки маршрутов
	appRouter := router.NewRouter(nearbyHandler, popularHandler, placeHandler)
	engine := appRouter.SetupRoutes()
	log.Println("✓ Routes configured")

	// === ДЕМОНСТРАЦИЯ ПРИНЦИПОВ ЧИСТОЙ АРХИТЕКТУРЫ ===
	/*
		1. Dependency Inversion Principle (DIP):
		   - Сервисы зависят от интерфейсов репозиториев
		   - Хендлеры зависят от интерфейсов сервисов
		   - Высокоуровневые модули не зависят от низкоуровневых

		2. Single Responsibility Principle (SRP):
		   - Каждый репозиторий отвечает только за свою сущность
		   - Каждый сервис содержит только бизнес-логику для своей области
		   - Каждый хендлер обрабатывает только свои HTTP запросы

		3. Open/Closed Principle (OCP):
		   - Легко добавить новый репозиторий (например, PostgreSQL)
		   - Легко добавить новый сервис для новой бизнес-логики
		   - Легко добавить новый хендлер для новых endpoints

		4. Interface Segregation Principle (ISP):
		   - Интерфейсы специфичны и небольшие
		   - NearbyService, PopularService, PlaceService - отдельные интерфейсы
		   - Каждый интерфейс содержит только нужные методы

		5. Liskov Substitution Principle (LSP):
		   - Любая реализация интерфейса может быть заменена на другую
		   - Mock репозитории легко заменяются на реальные
	*/

	// === ИНФОРМАЦИЯ О ДОСТУПНЫХ ENDPOINTS ===
	log.Println("🚀 Server starting on :8080")
	log.Println("📚 Available endpoints:")
	log.Println("  GET /api/v1/nearby?lat=55.7&lng=37.6  - поиск ближайших мест")
	log.Println("  GET /api/v1/popular                   - популярные места")
	log.Println("  GET /api/v1/places/:id                - место по ID")
	log.Println("  GET /api/v1/places                    - все места")
	log.Println("  GET /api/v1/health                    - проверка состояния сервиса")
	log.Println("📖 Backward compatibility endpoints:")
	log.Println("  GET /nearby?lat=55.7&lng=37.6         - поиск ближайших мест")
	log.Println("  GET /popular                          - популярные места")
	log.Println("  GET /health                           - проверка состояния")

	// === ДЕМОНСТРАЦИЯ ЛЕГКОСТИ РАСШИРЕНИЯ ===
	/*
		Чтобы добавить новый функционал:

		1. Добавить новый репозиторий:
		   - Создать файл review_repository.go
		   - Реализовать интерфейс ReviewRepository
		   - Добавить в main.go: reviewRepo := repository.NewReviewRepository()

		2. Добавить новый сервис:
		   - Создать файл review_service.go
		   - Реализовать интерфейс ReviewService
		   - Добавить в main.go: reviewService := service.NewReviewService(reviewRepo)

		3. Добавить новый хендлер:
		   - Создать файл review_handler.go
		   - Реализовать методы для обработки HTTP запросов
		   - Добавить в main.go: reviewHandler := handler.NewReviewHandler(reviewService)

		4. Добавить новые маршруты:
		   - Обновить router.go
		   - Добавить новые endpoints
	*/

	// Запуск HTTP сервера
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
