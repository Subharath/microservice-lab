package repository

import (
	"errors"
	"order-service/models"
	"strconv"
	"sync"
	"time"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

// OrderRepository handles data access for orders
type OrderRepository struct {
	orders    map[string]*models.Order
	currentID int
	mu        sync.RWMutex
}

// NewOrderRepository creates a new order repository with sample data
func NewOrderRepository() *OrderRepository {
	repo := &OrderRepository{
		orders:    make(map[string]*models.Order),
		currentID: 1,
	}

	// Add sample data
	repo.Create(&models.CreateOrderRequest{
		CustomerName:  "John Doe",
		CustomerEmail: "john@example.com",
		Items: []models.OrderItem{
			{ItemID: "1", Name: "Laptop", Quantity: 1, Price: 1299.99},
		},
	})

	repo.Create(&models.CreateOrderRequest{
		CustomerName:  "Jane Smith",
		CustomerEmail: "jane@example.com",
		Items: []models.OrderItem{
			{ItemID: "2", Name: "Wireless Mouse", Quantity: 2, Price: 29.99},
			{ItemID: "3", Name: "USB-C Cable", Quantity: 3, Price: 14.99},
		},
	})

	return repo
}

// FindAll returns all orders
func (r *OrderRepository) FindAll() []*models.Order {
	r.mu.RLock()
	defer r.mu.RUnlock()

	orders := make([]*models.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders
}

// FindByID returns an order by ID
func (r *OrderRepository) FindByID(id string) (*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, ErrOrderNotFound
	}
	return order, nil
}

// Create creates a new order
func (r *OrderRepository) Create(req *models.CreateOrderRequest) *models.Order {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Calculate total amount
	var total float64
	for _, item := range req.Items {
		total += item.Price * float64(item.Quantity)
	}

	order := &models.Order{
		ID:            strconv.Itoa(r.currentID),
		CustomerName:  req.CustomerName,
		CustomerEmail: req.CustomerEmail,
		Items:         req.Items,
		TotalAmount:   total,
		Status:        models.StatusPending,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	r.orders[order.ID] = order
	r.currentID++

	return order
}

// Update updates an existing order
func (r *OrderRepository) Update(id string, req *models.UpdateOrderRequest) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, ErrOrderNotFound
	}

	if req.CustomerName != "" {
		order.CustomerName = req.CustomerName
	}
	if req.CustomerEmail != "" {
		order.CustomerEmail = req.CustomerEmail
	}
	if req.Status != "" {
		order.Status = req.Status
	}

	order.UpdatedAt = time.Now()
	return order, nil
}

// Delete removes an order
func (r *OrderRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[id]; !exists {
		return ErrOrderNotFound
	}

	delete(r.orders, id)
	return nil
}

// UpdateStatus updates the status of an order
func (r *OrderRepository) UpdateStatus(id string, status models.OrderStatus) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, ErrOrderNotFound
	}

	order.Status = status
	order.UpdatedAt = time.Now()
	return order, nil
}
