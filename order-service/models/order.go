package models

import (
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivered OrderStatus = "delivered"
	StatusCancelled OrderStatus = "cancelled"
)

// OrderItem represents an item within an order
type OrderItem struct {
	ItemID   string  `json:"itemId" binding:"required"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity" binding:"required,min=1"`
	Price    float64 `json:"price"`
}

// Order represents a customer order
type Order struct {
	ID            string      `json:"id"`
	CustomerName  string      `json:"customerName" binding:"required"`
	CustomerEmail string      `json:"customerEmail" binding:"required,email"`
	Items         []OrderItem `json:"items" binding:"required,min=1,dive"`
	TotalAmount   float64     `json:"totalAmount"`
	Status        OrderStatus `json:"status"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	CustomerName  string      `json:"customerName" binding:"required"`
	CustomerEmail string      `json:"customerEmail" binding:"required,email"`
	Items         []OrderItem `json:"items" binding:"required,min=1,dive"`
}

// UpdateOrderRequest represents the request body for updating an order
type UpdateOrderRequest struct {
	CustomerName  string      `json:"customerName,omitempty"`
	CustomerEmail string      `json:"customerEmail,omitempty"`
	Status        OrderStatus `json:"status,omitempty"`
}

// OrderResponse wraps the order response
type OrderResponse struct {
	Success bool   `json:"success"`
	Data    *Order `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// OrderListResponse wraps multiple orders response
type OrderListResponse struct {
	Success bool     `json:"success"`
	Count   int      `json:"count"`
	Data    []*Order `json:"data"`
}
