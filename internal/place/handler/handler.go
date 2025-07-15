package handler

import (
	"net/http"
	"place_service/internal/place/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Создаёт HTTP-обработчик для поиска ближайших мест.
// Получает координаты из query-параметров, вызывает сервис и возвращает результат в формате JSON.
func NearbyHandler(svc service.NearbyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
		lng, _ := strconv.ParseFloat(c.Query("lng"), 64)

		places, err := svc.FindNearby(lat, lng)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, places)
	}
}

// Создаёт HTTP-обработчик для получения популярных мест.
// Просто вызывает сервис и возвращает результат в формате JSON.
func PopularHandler(svc service.PopularService) gin.HandlerFunc {
	return func(c *gin.Context) {
		places, err := svc.GetPopular()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, places)
	}
}
