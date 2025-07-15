package repository

import "place_service/internal/place/models"

type MockNearbyRepo struct{}

func NewMockNearbyRepo() *MockNearbyRepo {
	return &MockNearbyRepo{}
}

func (r *MockNearbyRepo) GetNearbyPlaces(lat, lng float64) ([]models.Place, error) {
	return []models.Place{
		{ID: 1, Name: "Near Park", Lat: lat, Lng: lng},
	}, nil
}
