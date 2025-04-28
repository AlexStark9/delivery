package courier

import (
	"delivery/internal/pkg/errs"
	"github.com/google/uuid"
)

// StoragePlace represents Entity with an ID, name, total volume, and an optional order ID.
type StoragePlace struct {
	id          uuid.UUID
	name        string
	totalVolume int
	order       *uuid.UUID
}

// NewStoragePlace creates a new StoragePlace instance with the given ID, name, total volume, and optional order ID.
func NewStoragePlace(name string, volume int) (*StoragePlace, error) {
	if name == "" {
		return nil, errs.NewValueIsRequiredError("name")
	}
	if volume <= 0 {
		return nil, errs.NewValueIsRequiredError("volume")
	}
	return &StoragePlace{
		id:          uuid.New(),
		name:        name,
		totalVolume: volume,
	}, nil
}

// GetID returns the ID of the StoragePlace.
func (sp *StoragePlace) GetID() uuid.UUID {
	return sp.id
}

// GetName returns the name of the StoragePlace.
func (sp *StoragePlace) GetName() string {
	return sp.name
}

// GetTotalVolume returns the total volume of the StoragePlace.
func (sp *StoragePlace) GetTotalVolume() int {
	return sp.totalVolume
}

// GetOrderID returns the order ID of the StoragePlace, if it exists.
func (sp *StoragePlace) GetOrderID() *uuid.UUID {
	return sp.order
}

// Equal checks if two StoragePlace instances are equal based on their ID.
func (sp *StoragePlace) Equal(other *StoragePlace) bool {
	if other == nil {
		return false
	}
	return sp.id == other.id
}

// AvailableVolume calculates the available volume in the StoragePlace.
func (sp *StoragePlace) AvailableVolume(volume int) (bool, error) {
	if volume <= 0 {
		return false, errs.NewValueIsRequiredError("volume cannot be less than or equal to 0")
	}
	if sp.IsEmptyOrder() {
		return false, nil
	}
	return volume <= sp.totalVolume, nil
}

// StoreOrder associates an order with the StoragePlace.
func (sp *StoragePlace) StoreOrder(orderID uuid.UUID, volume int) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID cannot be empty")
	}
	if volume <= 0 {
		return errs.NewValueIsRequiredError("volume cannot be less than or equal to 0")
	}

	storeOrder, err := sp.AvailableVolume(volume)
	if err != nil {
		return err
	}

	if !storeOrder {
		return errs.NewValueIsRequiredError("volume is more than permissible")
	}

	sp.order = &orderID
	return nil
}

// RemoveOrder disassociates the order from the StoragePlace.
func (sp *StoragePlace) RemoveOrder(orderID uuid.UUID) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID")
	}
	// Check if the orderID matches the one stored in the StoragePlace
	if sp.order == nil || *sp.order != orderID {
		return errs.NewValueIsRequiredError("orderID does not match")
	}
	// Remove the order by setting it to nil
	sp.order = nil
	return nil
}

// IsEmptyOrder checks if the StoragePlace is empty.
func (sp *StoragePlace) IsEmptyOrder() bool {
	return sp.order != nil
}
