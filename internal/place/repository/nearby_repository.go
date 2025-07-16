package repository

import (
	"place_service/internal/place/models"
)

// NearbyRepositoryImpl реализация репозитория для ближайших мест
type NearbyRepositoryImpl struct {
	places []models.Place
}

// NewNearbyRepository создает новый репозиторий для ближайших мест
func NewNearbyRepository() NearbyRepository {
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
		{
			ID:    3,
			Name:  "Третьяковская галерея",
			Lat:   55.7414,
			Lng:   37.6207,
			Photo: "https://example.com/tretyakov.jpg",
		},
		{
			ID:    4,
			Name:  "Воробьевы горы",
			Lat:   55.7105,
			Lng:   37.5439,
			Photo: "https://example.com/sparrow-hills.jpg",
		},
		{
			ID:    5,
			Name:  "Большой театр",
			Lat:   55.7596,
			Lng:   37.6189,
			Photo: "https://example.com/bolshoi.jpg",
		},
	}

	return &NearbyRepositoryImpl{
		places: places,
	}
}

// GetNearbyPlaces возвращает места рядом с указанными координатами
func (r *NearbyRepositoryImpl) GetNearbyPlaces(lat, lng float64) ([]models.Place, error) {
	// Простая логика поиска ближайших мест
	nearbyPlaces := make([]models.Place, 0)

	// Находим места в радиусе ~0.1 градуса (примерно 10 км)
	for _, place := range r.places {
		distance := place.DistanceTo(lat, lng)
		if distance < 0.01 { // Примерно 1 км в градусах
			nearbyPlaces = append(nearbyPlaces, place)
		}
	}

	// Если ничего не найдено рядом, возвращаем первые 3 места
	if len(nearbyPlaces) == 0 {
		if len(r.places) >= 3 {
			nearbyPlaces = r.places[:3]
		} else {
			nearbyPlaces = r.places
		}
	}

	return nearbyPlaces, nil
}
