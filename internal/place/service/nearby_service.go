package service

import (
	"place_service/internal/place/models"
	"place_service/internal/place/repository"
)

// NearbyServiceImpl реализация сервиса для поиска ближайших мест
type NearbyServiceImpl struct {
	nearbyRepo repository.NearbyRepository
}

// NewNearbyService создает новый сервис для поиска ближайших мест
func NewNearbyService(repo repository.NearbyRepository) NearbyService {
	return &NearbyServiceImpl{
		nearbyRepo: repo,
	}
}

// FindNearby находит места рядом с указанными координатами
func (s *NearbyServiceImpl) FindNearby(lat, lng float64) ([]models.Place, error) {
	// Валидация входных данных
	if lat < -90 || lat > 90 {
		return nil, NewValidationError("latitude must be between -90 and 90")
	}
	if lng < -180 || lng > 180 {
		return nil, NewValidationError("longitude must be between -180 and 180")
	}

	// Получение данных из репозитория
	places, err := s.nearbyRepo.GetNearbyPlaces(lat, lng)
	if err != nil {
		return nil, NewRepositoryError("failed to get nearby places", err)
	}

	// Фильтрация только валидных мест
	validPlaces := make([]models.Place, 0)
	for _, place := range places {
		if place.IsValid() {
			validPlaces = append(validPlaces, place)
		}
	}

	return validPlaces, nil
}
