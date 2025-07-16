package service

import "place_service/internal/place/models"

// NearbyService интерфейс для сервиса поиска ближайших мест
type NearbyService interface {
	FindNearby(lat, lng float64) ([]models.Place, error)
}

// PopularService интерфейс для сервиса популярных мест
type PopularService interface {
	GetPopular() ([]models.Place, error)
}

// PlaceService интерфейс для сервиса работы с местами по ID
type PlaceService interface {
	GetPlaceByID(id int) (*models.Place, error)
	GetAllPlaces() ([]models.Place, error)
	AddPlace(place models.Place) error
}
