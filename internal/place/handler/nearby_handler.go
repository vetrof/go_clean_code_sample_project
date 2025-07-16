package handler

import (
	"net/http"
	"place_service/internal/place/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NearbyHandler хендлер для работы с ближайшими местами
type NearbyHandler struct {
	nearbyService service.NearbyServiceInterface
}

// NewNearbyHandler создает новый хендлер для ближайших мест
func NewNearbyHandler(service service.NearbyServiceInterface) *NearbyHandler {
	return &NearbyHandler{
		nearbyService: service,
	}
}

// GetNearbyPlaces обработчик для GET /nearby?lat=55.7&lng=37.6
func (h *NearbyHandler) GetNearbyPlaces(c *gin.Context) {
	// Парсинг параметров запроса
	latStr := c.Query("lat")
	lngStr := c.Query("lng")

	// Валидация наличия параметров
	if latStr == "" || lngStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "lat and lng parameters are required",
			"code":  "missing_parameters",
		})
		return
	}

	// Парсинг координат
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid latitude format",
			"code":  "invalid_latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid longitude format",
			"code":  "invalid_longitude",
		})
		return
	}

	// Вызов сервиса
	places, err := h.nearbyService.FindNearby(lat, lng)
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
