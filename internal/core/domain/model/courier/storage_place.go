package courier

import (
	"delivery/internal/pkg/errs"
	"github.com/google/uuid"
)

// StoragePlace represents Entity with an ID, name, total volume, and an optional orderID ID.
type StoragePlace struct {
	id          uuid.UUID
	name        string
	totalVolume int
	orderID     *uuid.UUID
}

// NewStoragePlace creates a new StoragePlace instance with the given ID, name, total volume, and optional orderID ID.
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

// ID returns the ID of the StoragePlace.
func (sp *StoragePlace) ID() uuid.UUID {
	return sp.id
}

// Name returns the name of the StoragePlace.
func (sp *StoragePlace) Name() string {
	return sp.name
}

// TotalVolume returns the total volume of the StoragePlace.
func (sp *StoragePlace) TotalVolume() int {
	return sp.totalVolume
}

// OrderID returns the orderID ID of the StoragePlace, if it exists.
func (sp *StoragePlace) OrderID() *uuid.UUID {
	return sp.orderID
}

// Equal checks if two StoragePlace instances are equal based on their ID.
func (sp *StoragePlace) Equal(other *StoragePlace) bool {
	if other == nil {
		return false
	}
	return sp.id == other.id
}

// CanCurrentlyStore calculates the available volume in the StoragePlace.
func (sp *StoragePlace) CanCurrentlyStore(volume int) (bool, error) {
	if volume <= 0 {
		return false, errs.NewValueIsRequiredError("volume cannot be less than or equal to 0")
	}
	if sp.hasOrder() {
		return false, nil
	}
	return volume <= sp.totalVolume, nil
}

// StoreOrder associates an orderID with the StoragePlace.
func (sp *StoragePlace) StoreOrder(orderID uuid.UUID, volume int) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID cannot be empty")
	}
	if volume <= 0 {
		return errs.NewValueIsRequiredError("volume cannot be less than or equal to 0")
	}

	storeOrder, err := sp.CanCurrentlyStore(volume)
	if err != nil {
		return err
	}

	if !storeOrder {
		return errs.NewValueIsRequiredError("volume is more than permissible")
	}

	sp.orderID = &orderID
	return nil
}

// RemoveOrder disassociates the orderID from the StoragePlace.
func (sp *StoragePlace) RemoveOrder(orderID uuid.UUID) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID")
	}
	// Check if the orderID matches the one stored in the StoragePlace
	if sp.orderID == nil || *sp.orderID != orderID {
		return errs.NewValueIsRequiredError("orderID does not match")
	}
	// Remove the orderID by setting it to nil
	sp.orderID = nil
	return nil
}

// hasOrder checks if the StoragePlace is empty.
func (sp *StoragePlace) hasOrder() bool {
	return sp.orderID != nil
}
