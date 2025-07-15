package service

import (
	"place_service/internal/place/models"
)

type NearbyService interface {
	FindNearby(lat, lng float64) ([]models.Place, error)
}

type nearbyService struct {
	repo NearbyRepository
}

func NewNearbyService(repo NearbyRepository) NearbyService {
	return &nearbyService{repo}
}

func (s *nearbyService) FindNearby(lat, lng float64) ([]models.Place, error) {
	return s.repo.GetNearbyPlaces(lat, lng)
}

// Интерфейс репозитория — тоже в этом же файле (тонкий)
type NearbyRepository interface {
	GetNearbyPlaces(lat, lng float64) ([]models.Place, error)
}
