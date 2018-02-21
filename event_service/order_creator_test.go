package event_service

import (
	"testing"

	"strings"

	"github.com/timhugh/ticket-engine/lib/repos"
)

type MockOrderRepository struct {
	Order *repos.Order
}

func (m *MockOrderRepository) Store(order repos.Order) {
	m.Order = &order
}

func (m *MockOrderRepository) Find(id string) *repos.Order {
	return m.Order
}

func TestRejectsDuplicates(t *testing.T) {
	orderCreator := OrderCreator{
		&MockOrderRepository{
			Order: &repos.Order{},
		},
	}

	err := orderCreator.Create("order_id", "location_id")
	if err == nil {
		t.Error("Expected duplicate order to fail creating.")
	}
	if !strings.Contains(err.Error(), "duplicate order") {
		t.Errorf("Expected error to include 'duplicate order' but received %s", err.Error())
	}
}

func TestStoresValidOrder(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	orderCreator := OrderCreator{
		OrderRepository: mockRepo,
	}

	orderID := "order_id"
	locationID := "location_id"

	if err := orderCreator.Create(orderID, locationID); err != nil {
		t.Errorf("Expected successful creation of order but received error: %s", err.Error())
	}

	order := mockRepo.Order
	if order.ID != orderID {
		t.Errorf("Expected to create an order with ID %s but found %s", orderID, order.ID)
	}
	if order.LocationID != locationID {
		t.Errorf("Expected to create an order with locationID %s but found %s", locationID, order.LocationID)
	}
}