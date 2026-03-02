package routes

import (
	"net/http"
	"time"

	"order-service/config"
	"order-service/controllers"
	"order-service/middleware"
	"order-service/repository"
	"order-service/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router
func SetupRouter(cfg *config.Config) *gin.Engine {
	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Global middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger(cfg.ServiceName))
	router.Use(middleware.CORS(middleware.CORSConfig{
		AllowOrigin: cfg.CORSOrigin,
	}))

	// Initialize dependencies
	orderRepo := repository.NewOrderRepository()
	itemClient := services.NewItemClient(cfg.ItemServiceURL)
	orderService := services.NewOrderService(orderRepo, itemClient)
	orderController := controllers.NewOrderController(orderService)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "OK",
			"service":     cfg.ServiceName,
			"version":     cfg.Version,
			"environment": cfg.Environment,
			"timestamp":   time.Now().Format(time.RFC3339),
		})
	})

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to " + cfg.ServiceName + " API",
			"version": cfg.Version,
			"endpoints": gin.H{
				"health": "/health",
				"orders": "/api/orders",
			},
		})
	})

	// API v1 routes
	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.GET("", orderController.GetAllOrders)
			orders.GET("/:id", orderController.GetOrderByID)
			orders.POST("", orderController.CreateOrder)
			orders.PUT("/:id", orderController.UpdateOrder)
			orders.DELETE("/:id", orderController.DeleteOrder)
			orders.PATCH("/:id/status", orderController.UpdateOrderStatus)
		}
	}

	return router
}
