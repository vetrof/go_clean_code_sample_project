package repository

import "place_service/internal/place/models"

type MockPopularRepo struct{}

func NewMockPopularRepo() *MockPopularRepo {
	return &MockPopularRepo{}
}

func (r *MockPopularRepo) GetPopularPlaces() ([]models.Place, error) {
	return []models.Place{
		{ID: 2, Name: "Popular Museum"},
	}, nil
}
