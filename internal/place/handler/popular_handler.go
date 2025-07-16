package handler

import (
	"net/http"
	"place_service/internal/place/service"

	"github.com/gin-gonic/gin"
)

// PopularHandler хендлер для работы с популярными местами
type PopularHandler struct {
	popularService service.PopularService
}

// NewPopularHandler создает новый хендлер для популярных мест
func NewPopularHandler(service service.PopularService) *PopularHandler {
	return &PopularHandler{
		popularService: service,
	}
}

// GetPopularPlaces обработчик для GET /popular
func (h *PopularHandler) GetPopularPlaces(c *gin.Context) {
	// Вызов сервиса
	places, err := h.popularService.GetPopular()
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
