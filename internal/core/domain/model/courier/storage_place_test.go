package courier

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCreateValidStoragePlace tests the creation of a valid StoragePlace instance.
func TestCreateValidStoragePlace(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)

	assert.NoError(t, err)
	assert.NotEmpty(t, sp)
	assert.Equal(t, "Test Place", sp.GetName())
	assert.Equal(t, 100, sp.GetTotalVolume())
}

// TestCreateInvalidStoragePlace tests the creation of an invalid StoragePlace instance.
func TestCreateInvalidStoragePlace(t *testing.T) {
	t.Parallel()
	// Test case for empty name
	testData := []struct {
		name   string
		volume int
	}{
		// Test case for empty name
		{"", 100},
		// Test case for zero volume
		{"Test Place", 0},
	}
	// Iterate over test cases
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			_, err := NewStoragePlace(data.name, data.volume)
			assert.Error(t, err)
		})
	}
}

// TestEqualStoragePlace tests the equality of two StoragePlace instances.
func TestEqualStoragePlace(t *testing.T) {
	t.Parallel()

	sp1, err := NewStoragePlace("Test Place 1", 100)
	assert.NoError(t, err)

	sp2, err := NewStoragePlace("Test Place 2", 100)
	assert.NoError(t, err)

	// Test equality
	assert.True(t, sp1.Equal(sp1))
	assert.False(t, sp1.Equal(sp2))
}

// TestAvailableVolumeValid tests the available volume calculation with valid values.
func TestAvailableVolumeValid(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)

	// Test case for valid volume
	available, err := sp.AvailableVolume(50)
	assert.NoError(t, err)
	assert.True(t, available)
}

// TestAvailableVolumeReturnError tests the available volume calculation with invalid values.
func TestAvailableVolumeReturnError(t *testing.T) {
	t.Parallel()
	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)

	// Test case for volume exceeding total volume
	available, err := sp.AvailableVolume(150)
	assert.NoError(t, err)
	assert.False(t, available)
}

// TestAvailableVolumeInvalidValue tests the available volume calculation with invalid values.
func TestAvailableVolumeInvalidValue(t *testing.T) {
	t.Parallel()
	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)

	// Test case for invalid volume
	available, err := sp.AvailableVolume(0)
	assert.Error(t, err)
	assert.False(t, available)
}

// TestStorageOrderWithInvalidVolume tests the storage order with invalid volume.
func TestStorageOrderWithInvalidVolume(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)

	// Test case for invalid volume
	err = sp.StoreOrder(uuid.New(), 0)
	assert.Error(t, err)
}

// TestStoreOrderFailed tests the storage order with failed conditions.
func TestStoreOrderFailed(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)
	err = sp.StoreOrder(uuid.New(), 100)
	assert.NoError(t, err)

	// Test case for storing order with empty order ID
	err = sp.StoreOrder(uuid.New(), 50)
	assert.Error(t, err)
}

// TestRemoveOrder tests the removal of an order from the StoragePlace.
func TestRemoveOrder(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)
	// Store an order
	orderID := uuid.New()
	err = sp.StoreOrder(orderID, 50)
	assert.NoError(t, err)
	// Remove the order
	err = sp.RemoveOrder(orderID)
	assert.NoError(t, err)
	assert.Nil(t, sp.GetOrderID())
}

// TestRemoveOrderWithInvalidID tests the removal of an order with an invalid ID.
func TestRemoveOrderWithInvalidID(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)
	err = sp.StoreOrder(uuid.New(), 100)
	assert.NoError(t, err)

	// Test case for removing order with invalid ID
	err = sp.RemoveOrder(uuid.New())
	assert.Error(t, err)
}

// TestRemoveOrderWithEmptyID tests the removal of an order with an empty ID.
func TestRemoveOrderWithEmptyID(t *testing.T) {
	t.Parallel()

	sp, err := NewStoragePlace("Test Place", 100)
	assert.NoError(t, err)

	// Test case for removing order with empty ID
	err = sp.RemoveOrder(uuid.Nil)
	assert.Error(t, err)
}
