package service

import (
	"place_service/internal/place/models"
	"place_service/internal/place/repository"
)

// PopularServiceInterface интерфейс для сервиса популярных мест
type PopularServiceInterface interface {
	GetPopular() ([]models.Place, error)
}

// PopularServiceImpl реализация сервиса для популярных мест
type PopularServiceImpl struct {
	popularRepo repository.PopularRepositoryInterface
}

// NewPopularService создает новый сервис для популярных мест
func NewPopularService(repo repository.PopularRepositoryInterface) PopularServiceInterface {
	return &PopularServiceImpl{
		popularRepo: repo,
	}
}

// GetPopular получает список популярных мест
func (s *PopularServiceImpl) GetPopular() ([]models.Place, error) {
	// Получение всех мест из репозитория
	allPlaces, err := s.popularRepo.GetPopularPlaces()
	if err != nil {
		return nil, NewRepositoryError("failed to get places from repository", err)
	}

	// БИЗНЕС-ЛОГИКА: фильтрация только валидных мест
	validPlaces := make([]models.Place, 0)
	for _, place := range allPlaces {
		if place.IsValid() {
			validPlaces = append(validPlaces, place)
		}
	}

	// БИЗНЕС-ПРАВИЛО: популярными считаются места с определенными ID
	// В реальном приложении здесь была бы логика на основе рейтингов, отзывов и т.д.
	popularPlaces := make([]models.Place, 0)
	popularIDs := map[int]bool{1: true, 2: true, 3: true, 5: true, 8: true} // Популярные ID

	for _, place := range validPlaces {
		if popularIDs[place.ID] {
			popularPlaces = append(popularPlaces, place)
		}
	}

	// БИЗНЕС-ПРАВИЛО: максимум 5 популярных мест
	if len(popularPlaces) > 5 {
		popularPlaces = popularPlaces[:5]
	}

	return popularPlaces, nil
}
