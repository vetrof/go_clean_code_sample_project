package service

import (
	"place_service/internal/place/models"
	"place_service/internal/place/repository"
)

// NearbyServiceInterface описывает интерфейс для поиска ближайших мест
type NearbyServiceInterface interface {
	FindNearby(lat, lng float64) ([]models.Place, error)
}

// NearbyServiceImpl — реализация сервиса
type NearbyServiceImpl struct {
	placeRepo repository.NearbyRepositoryInterface
}

// NewNearbyService создает новый экземпляр сервиса
func NewNearbyService(repo repository.NearbyRepositoryInterface) NearbyServiceInterface {
	return &NearbyServiceImpl{
		placeRepo: repo,
	}
}

// FindNearby получает список ближайших мест
func (s *NearbyServiceImpl) FindNearby(lat, lng float64) ([]models.Place, error) {
	places, err := s.placeRepo.GetNearbyPlaces(lat, lng)
	if err != nil {
		return nil, err
	}

	// Пока что возвращаем все места — фильтрация может быть добавлена здесь
	return places, nil
}
