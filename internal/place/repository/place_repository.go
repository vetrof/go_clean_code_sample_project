package repository

import (
	"fmt"
	"place_service/internal/place/models"
)

// PlaceRepositoryImpl реализация репозитория для работы с местами по ID
type PlaceRepositoryImpl struct {
	places []models.Place
}

// NewPlaceRepository создает новый репозиторий для работы с местами
func NewPlaceRepository() PlaceRepository {
	// Создаем набор тестовых данных для всех мест
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
		{
			ID:    6,
			Name:  "Московский зоопарк",
			Lat:   55.7558,
			Lng:   37.6142,
			Photo: "https://example.com/zoo.jpg",
		},
		{
			ID:    7,
			Name:  "Патриаршие пруды",
			Lat:   55.7649,
			Lng:   37.5947,
			Photo: "https://example.com/patriarshie.jpg",
		},
		{
			ID:    8,
			Name:  "Музей Пушкина",
			Lat:   55.7474,
			Lng:   37.6049,
			Photo: "https://example.com/pushkin-museum.jpg",
		},
	}

	return &PlaceRepositoryImpl{
		places: places,
	}
}

// GetPlaceByID возвращает место по ID
func (r *PlaceRepositoryImpl) GetPlaceByID(id int) (*models.Place, error) {
	// Валидация ID
	if id <= 0 {
		return nil, fmt.Errorf("invalid id: %d", id)
	}

	// Поиск места по ID
	for _, place := range r.places {
		if place.ID == id {
			return &place, nil
		}
	}

	// Место не найдено
	return nil, fmt.Errorf("place with id %d not found", id)
}

// GetAllPlaces возвращает все места
func (r *PlaceRepositoryImpl) GetAllPlaces() ([]models.Place, error) {
	return r.places, nil
}

// AddPlace добавляет новое место (для демонстрации расширяемости)
func (r *PlaceRepositoryImpl) AddPlace(place models.Place) error {
	// Проверка на дубликат ID
	for _, existingPlace := range r.places {
		if existingPlace.ID == place.ID {
			return fmt.Errorf("place with id %d already exists", place.ID)
		}
	}

	// Валидация места
	if !place.IsValid() {
		return fmt.Errorf("invalid place data")
	}

	r.places = append(r.places, place)
	return nil
}

// Count возвращает количество мест в репозитории
func (r *PlaceRepositoryImpl) Count() int {
	return len(r.places)
}
