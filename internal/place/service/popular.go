package service

import "place_service/internal/place/models"

type PopularService interface {
	GetPopular() ([]models.Place, error)
}

type popularService struct {
	repo PopularRepository
}

func NewPopularService(repo PopularRepository) PopularService {
	return &popularService{repo}
}

func (s *popularService) GetPopular() ([]models.Place, error) {
	return s.repo.GetPopularPlaces()
}

type PopularRepository interface {
	GetPopularPlaces() ([]models.Place, error)
}
