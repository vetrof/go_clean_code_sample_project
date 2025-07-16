package repository

import (
	"place_service/internal/place/models"
)

// PopularRepositoryImpl реализация репозитория для популярных мест
type PopularRepositoryImpl struct {
	places []models.Place
}

// NewPopularRepository создает новый репозиторий для популярных мест
func NewPopularRepository() PopularRepository {
	// Создаем набор тестовых данных для популярных мест
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
			Name:  "Большой театр",
			Lat:   55.7596,
			Lng:   37.6189,
			Photo: "https://example.com/bolshoi.jpg",
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
			Name:  "Московский зоопарк",
			Lat:   55.7558,
			Lng:   37.6142,
			Photo: "https://example.com/zoo.jpg",
		},
		{
			ID:    5,
			Name:  "Музей Пушкина",
			Lat:   55.7474,
			Lng:   37.6049,
			Photo: "https://example.com/pushkin-museum.jpg",
		},
	}

	return &PopularRepositoryImpl{
		places: places,
	}
}

// GetPopularPlaces возвращает список популярных мест
func (r *PopularRepositoryImpl) GetPopularPlaces() ([]models.Place, error) {
	// В реальном приложении здесь была бы логика сортировки по популярности
	// Возвращаем все популярные места
	return r.places, nil
}
