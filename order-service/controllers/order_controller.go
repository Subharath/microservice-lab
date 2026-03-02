package controllers

import (
	"errors"
	"net/http"
	"order-service/models"
	"order-service/repository"
	"order-service/services"

	"github.com/gin-gonic/gin"
)

// OrderController handles HTTP requests for orders
type OrderController struct {
	service *services.OrderService
}

// NewOrderController creates a new order controller
func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// GetAllOrders handles GET /api/orders
func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders := c.service.GetAllOrders()

	ctx.JSON(http.StatusOK, models.OrderListResponse{
		Success: true,
		Count:   len(orders),
		Data:    orders,
	})
}

// GetOrderByID handles GET /api/orders/:id
func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := c.service.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, models.OrderResponse{
				Success: false,
				Error:   "Order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.OrderResponse{
		Success: true,
		Data:    order,
	})
}

// CreateOrder handles POST /api/orders
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	order, err := c.service.CreateOrder(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.OrderResponse{
		Success: true,
		Data:    order,
	})
}

// UpdateOrder handles PUT /api/orders/:id
func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.UpdateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	order, err := c.service.UpdateOrder(id, &req)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, models.OrderResponse{
				Success: false,
				Error:   "Order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.OrderResponse{
		Success: true,
		Data:    order,
	})
}

// DeleteOrder handles DELETE /api/orders/:id
func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DeleteOrder(id)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order deleted successfully",
	})
}

// UpdateOrderStatus handles PATCH /api/orders/:id/status
func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	var req struct {
		Status models.OrderStatus `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	order, err := c.service.UpdateOrderStatus(id, req.Status)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, models.OrderResponse{
				Success: false,
				Error:   "Order not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, models.OrderResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.OrderResponse{
		Success: true,
		Data:    order,
	})
}
