package service

import (
	"place_service/internal/place/models"
	"place_service/internal/place/repository"
)

// PopularServiceImpl реализация сервиса для популярных мест
type PopularServiceImpl struct {
	popularRepo repository.PopularRepository
}

// NewPopularService создает новый сервис для популярных мест
func NewPopularService(repo repository.PopularRepository) PopularService {
	return &PopularServiceImpl{
		popularRepo: repo,
	}
}

// GetPopular получает список популярных мест
func (s *PopularServiceImpl) GetPopular() ([]models.Place, error) {
	// Получение данных из репозитория
	places, err := s.popularRepo.GetPopularPlaces()
	if err != nil {
		return nil, NewRepositoryError("failed to get popular places", err)
	}

	// Фильтрация только валидных мест
	validPlaces := make([]models.Place, 0)
	for _, place := range places {
		if place.IsValid() {
			validPlaces = append(validPlaces, place)
		}
	}

	// Бизнес-логика: сортировка по популярности (для демонстрации)
	// В реальном приложении здесь была бы более сложная логика
	if len(validPlaces) > 5 {
		validPlaces = validPlaces[:5] // Берем только топ-5 популярных
	}

	return validPlaces, nil
}
