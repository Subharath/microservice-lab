package services

import (
	"fmt"
	"log"
	"order-service/models"
	"order-service/repository"
)

// OrderService handles business logic for orders
type OrderService struct {
	repo       *repository.OrderRepository
	itemClient *ItemClient
}

// NewOrderService creates a new order service
func NewOrderService(repo *repository.OrderRepository, itemClient *ItemClient) *OrderService {
	return &OrderService{
		repo:       repo,
		itemClient: itemClient,
	}
}

// GetAllOrders returns all orders
func (s *OrderService) GetAllOrders() []*models.Order {
	return s.repo.FindAll()
}

// GetOrderByID returns an order by ID
func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	return s.repo.FindByID(id)
}

// CreateOrder creates a new order with item validation
func (s *OrderService) CreateOrder(req *models.CreateOrderRequest) (*models.Order, error) {
	// Prepare items for validation
	itemsToValidate := make([]struct {
		ItemID   string
		Quantity int
	}, len(req.Items))

	for i, item := range req.Items {
		itemsToValidate[i] = struct {
			ItemID   string
			Quantity int
		}{
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
	}

	// Validate items with item-service
	validatedItems, err := s.itemClient.ValidateItems(itemsToValidate)
	if err != nil {
		log.Printf("Item validation failed: %v", err)
		// For development, we'll allow orders even if item-service is down
		// In production, you might want to return the error
		log.Println("Warning: Proceeding without item validation")
	} else {
		// Enrich order items with details from item-service
		for i, orderItem := range req.Items {
			if item, exists := validatedItems[orderItem.ItemID]; exists {
				req.Items[i].Name = item.Name
				req.Items[i].Price = item.Price
			}
		}
	}

	// Create the order
	return s.repo.Create(req), nil
}

// UpdateOrder updates an existing order
func (s *OrderService) UpdateOrder(id string, req *models.UpdateOrderRequest) (*models.Order, error) {
	return s.repo.Update(id, req)
}

// DeleteOrder deletes an order
func (s *OrderService) DeleteOrder(id string) error {
	return s.repo.Delete(id)
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(id string, status models.OrderStatus) (*models.Order, error) {
	// Validate status transition
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if !s.isValidStatusTransition(order.Status, status) {
		return nil, fmt.Errorf("invalid status transition from %s to %s", order.Status, status)
	}

	return s.repo.UpdateStatus(id, status)
}

// isValidStatusTransition checks if a status transition is valid
func (s *OrderService) isValidStatusTransition(from, to models.OrderStatus) bool {
	validTransitions := map[models.OrderStatus][]models.OrderStatus{
		models.StatusPending:   {models.StatusConfirmed, models.StatusCancelled},
		models.StatusConfirmed: {models.StatusPaid, models.StatusCancelled},
		models.StatusPaid:      {models.StatusShipped, models.StatusCancelled},
		models.StatusShipped:   {models.StatusDelivered},
		models.StatusDelivered: {},
		models.StatusCancelled: {},
	}

	allowedTransitions, exists := validTransitions[from]
	if !exists {
		return false
	}

	for _, allowed := range allowedTransitions {
		if allowed == to {
			return true
		}
	}

	return false
}
