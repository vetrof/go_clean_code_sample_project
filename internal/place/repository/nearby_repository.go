package repository

import (
	"place_service/internal/place/models"
)

// NearbyRepositoryInterface интерфейс для работы с ближайшими местами
type NearbyRepositoryInterface interface {
	GetNearbyPlaces(lat, lng float64) ([]models.Place, error)
}

// NearbyRepositoryImpl реализация репозитория для ближайших мест
type NearbyRepositoryImpl struct {
	places []models.Place
}

// NewNearbyRepository создает новый репозиторий для ближайших мест
func MokNearbyRepository() NearbyRepositoryInterface {
	// Создаем набор тестовых данных для ближайших мест
	places := []models.Place{
		{
			ID:    1,
			Name:  "Красная площадь",
			Lat:   55.7539,
			Lng:   37.6208,
			Photo: "https://example.com/red-square.jpg",
		},
		{
			ID:    2,
			Name:  "Парк Горького",
			Lat:   55.7295,
			Lng:   37.6018,
			Photo: "https://example.com/gorky-park.jpg",
		},
	}

	return &NearbyRepositoryImpl{
		places: places,
	}
}

// GetNearbyPlaces возвращает все места (без фильтрации)
// Бизнес-логика фильтрации должна быть в сервисе
func (r *NearbyRepositoryImpl) GetNearbyPlaces(lat, lng float64) ([]models.Place, error) {
	// Репозиторий просто возвращает все доступные места
	// Фильтрация по расстоянию - это бизнес-логика, она должна быть в сервисе
	return r.places, nil
}
