package handler

import (
	"net/http"
	"place_service/internal/place/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PlaceHandler хендлер для работы с местами по ID
type PlaceHandler struct {
	placeService service.PlaceService
}

// NewPlaceHandler создает новый хендлер для работы с местами
func NewPlaceHandler(service service.PlaceService) *PlaceHandler {
	return &PlaceHandler{
		placeService: service,
	}
}

// GetPlaceByID обработчик для GET /places/:id
func (h *PlaceHandler) GetPlaceByID(c *gin.Context) {
	// Парсинг ID из пути
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
			"code":  "invalid_id",
		})
		return
	}

	// Вызов сервиса
	place, err := h.placeService.GetPlaceByID(id)
	if err != nil {
		// Обработка ошибок сервиса
		statusCode := service.GetErrorCode(err)
		errorType := service.GetErrorType(err)

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
			"code":  errorType,
		})
		return
	}

	// Проверка, что место найдено
	if place == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "place not found",
			"code":  "not_found",
		})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"place": place,
	})
}

// GetAllPlaces обработчик для GET /places
func (h *PlaceHandler) GetAllPlaces(c *gin.Context) {
	// Вызов сервиса
	places, err := h.placeService.GetAllPlaces()
	if err != nil {
		// Обработка ошибок сервиса
		statusCode := service.GetErrorCode(err)
		errorType := service.GetErrorType(err)

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
			"code":  errorType,
		})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"places": places,
		"count":  len(places),
	})
}

// HealthCheck обработчик для проверки состояния сервиса
func (h *PlaceHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "place-service",
		"version": "1.0.0",
	})
}
