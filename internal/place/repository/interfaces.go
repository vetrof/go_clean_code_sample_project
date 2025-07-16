package repository

import "place_service/internal/place/models"

// NearbyRepository интерфейс для работы с ближайшими местами
type NearbyRepository interface {
	GetNearbyPlaces(lat, lng float64) ([]models.Place, error)
}

// PopularRepository интерфейс для работы с популярными местами
type PopularRepository interface {
	GetPopularPlaces() ([]models.Place, error)
}

// PlaceRepository интерфейс для работы с местами по ID
type PlaceRepository interface {
	GetPlaceByID(id int) (*models.Place, error)
	GetAllPlaces() ([]models.Place, error)
	AddPlace(place models.Place) error
}
