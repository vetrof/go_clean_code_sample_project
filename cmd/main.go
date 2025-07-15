package main

import (
	// Импорт обработчиков, репозиториев, роутера и сервисов для работы с местами
	"place_service/internal/place/handler"
	"place_service/internal/place/repository"
	"place_service/internal/place/router"
	"place_service/internal/place/service"
)

// Точка входа в приложение
func main() {
	// Создание мок-репозиториев для поиска ближайших и популярных мест
	nearRepo := repository.NewMockNearbyRepo()
	popularRepo := repository.NewMockPopularRepo()

	// Создание сервисов, использующих соответствующие репозитории
	nearSvc := service.NewNearbyService(nearRepo)
	popSvc := service.NewPopularService(popularRepo)

	// Настройка роутера с обработчиками для поиска ближайших и популярных мест
	r := router.SetupRouter(
		handler.NearbyHandler(nearSvc),
		handler.PopularHandler(popSvc),
	)

	// Запуск HTTP-сервера на порту 8080
	r.Run(":8080")
}
