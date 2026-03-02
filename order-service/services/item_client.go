package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Item represents an item from the item-service
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// ItemResponse wraps the item service response
type ItemResponse struct {
	Success bool  `json:"success"`
	Data    *Item `json:"data"`
}

// ItemClient handles communication with the item-service
type ItemClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewItemClient creates a new item client
func NewItemClient(baseURL string) *ItemClient {
	return &ItemClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetItem fetches an item by ID from the item-service
func (c *ItemClient) GetItem(itemID string) (*Item, error) {
	url := fmt.Sprintf("%s/api/items/%s", c.baseURL, itemID)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to item-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("item %s not found", itemID)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("item-service returned status %d", resp.StatusCode)
	}

	var itemResp ItemResponse
	if err := json.NewDecoder(resp.Body).Decode(&itemResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !itemResp.Success {
		return nil, fmt.Errorf("item-service returned error")
	}

	return itemResp.Data, nil
}

// ValidateItems checks if all items exist and have sufficient stock
func (c *ItemClient) ValidateItems(items []struct {
	ItemID   string
	Quantity int
}) (map[string]*Item, error) {
	validatedItems := make(map[string]*Item)

	for _, orderItem := range items {
		item, err := c.GetItem(orderItem.ItemID)
		if err != nil {
			return nil, err
		}

		if item.Quantity < orderItem.Quantity {
			return nil, fmt.Errorf("insufficient stock for item %s: available %d, requested %d",
				item.Name, item.Quantity, orderItem.Quantity)
		}

		validatedItems[orderItem.ItemID] = item
	}

	return validatedItems, nil
}

// HealthCheck checks if item-service is healthy
func (c *ItemClient) HealthCheck() error {
	url := fmt.Sprintf("%s/health", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("item-service unreachable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("item-service unhealthy: status %d", resp.StatusCode)
	}

	return nil
}
