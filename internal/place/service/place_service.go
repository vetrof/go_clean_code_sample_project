package service

import (
	"place_service/internal/place/models"
	"place_service/internal/place/repository"
)

// PlaceServiceImpl реализация сервиса для работы с местами по ID
type PlaceServiceImpl struct {
	placeRepo repository.PlaceRepository
}

// NewPlaceService создает новый сервис для работы с местами
func NewPlaceService(repo repository.PlaceRepository) PlaceService {
	return &PlaceServiceImpl{
		placeRepo: repo,
	}
}

// GetPlaceByID получает место по ID
func (s *PlaceServiceImpl) GetPlaceByID(id int) (*models.Place, error) {
	// Валидация ID
	if id <= 0 {
		return nil, NewValidationError("id must be positive")
	}

	// Получение данных из репозитория
	place, err := s.placeRepo.GetPlaceByID(id)
	if err != nil {
		return nil, NewRepositoryError("failed to get place by id", err)
	}

	// Валидация полученного места
	if place != nil && !place.IsValid() {
		return nil, NewValidationError("place data is invalid")
	}

	return place, nil
}

// GetAllPlaces получает все места
func (s *PlaceServiceImpl) GetAllPlaces() ([]models.Place, error) {
	// Получение данных из репозитория
	places, err := s.placeRepo.GetAllPlaces()
	if err != nil {
		return nil, NewRepositoryError("failed to get all places", err)
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

// AddPlace добавляет новое место
func (s *PlaceServiceImpl) AddPlace(place models.Place) error {
	// Валидация места
	if !place.IsValid() {
		return NewValidationError("invalid place data")
	}

	// Проверка на дубликат ID через сервис
	existingPlace, err := s.placeRepo.GetPlaceByID(place.ID)
	if err == nil && existingPlace != nil {
		return NewValidationError("place with this id already exists")
	}

	// Добавление через репозиторий
	if err := s.placeRepo.AddPlace(place); err != nil {
		return NewRepositoryError("failed to add place", err)
	}

	return nil
}
