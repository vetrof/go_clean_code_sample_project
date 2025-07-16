package router

import (
	"place_service/internal/place/handler"

	"github.com/gin-gonic/gin"
)

// Router содержит все обработчики для настройки маршрутов
type Router struct {
	nearbyHandler  *handler.NearbyHandler
	popularHandler *handler.PopularHandler
	placeHandler   *handler.PlaceHandler
}

// NewRouter создает новый роутер с обработчиками
func NewRouter(
	nearbyHandler *handler.NearbyHandler,
	popularHandler *handler.PopularHandler,
	placeHandler *handler.PlaceHandler,
) *Router {
	return &Router{
		nearbyHandler:  nearbyHandler,
		popularHandler: popularHandler,
		placeHandler:   placeHandler,
	}
}

// SetupRoutes настраивает все маршруты для приложения
func (r *Router) SetupRoutes() *gin.Engine {
	// Создаем Gin engine
	engine := gin.Default()

	// Middleware для CORS
	engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Основные маршруты API
	api := engine.Group("/api/v1")
	{
		// Маршруты для работы с местами
		api.GET("/nearby", r.nearbyHandler.GetNearbyPlaces)
		api.GET("/popular", r.popularHandler.GetPopularPlaces)
		api.GET("/places/:id", r.placeHandler.GetPlaceByID)
		api.GET("/places", r.placeHandler.GetAllPlaces)

		// Служебные маршруты
		api.GET("/health", r.placeHandler.HealthCheck)
	}

	// Для обратной совместимости - старые маршруты
	engine.GET("/nearby", r.nearbyHandler.GetNearbyPlaces)
	engine.GET("/popular", r.popularHandler.GetPopularPlaces)
	engine.GET("/health", r.placeHandler.HealthCheck)

	return engine
}

// SetupRoutesWithMiddleware настраивает маршруты с дополнительными middleware
func (r *Router) SetupRoutesWithMiddleware(middlewares ...gin.HandlerFunc) *gin.Engine {
	engine := r.SetupRoutes()

	// Добавляем дополнительные middleware
	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	return engine
}
