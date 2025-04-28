package shared_kernel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCreateValidLocation tests the creation of a valid Location instance.
func TestCreateVal(t *testing.T) {
	t.Parallel()

	location, err := NewLocation(5, 5)
	assert.NoError(t, err)
	assert.NotEmpty(t, location)
	assert.Equal(t, 5, location.GetX())
	assert.Equal(t, 5, location.GetY())
}

// TestCreateInvalidLocation tests the creation of an invalid Location instance.
func TestCreateInvalidLocation(t *testing.T) {
	t.Parallel()

	_, err := NewLocation(11, -1)
	assert.Error(t, err)
}

// TestCalculateDistance tests the calculation of distance between two Location instances.
func TestCalculateDistance(t *testing.T) {
	t.Parallel()

	location1, err := NewLocation(5, 5)
	assert.NoError(t, err)
	location2, err := NewLocation(10, 10)
	assert.NoError(t, err)

	distance, err := location1.DistanceBetweenLocations(location2)
	assert.NoError(t, err)
	assert.Equal(t, 10, distance)
}

// TestReturnErrorWhenLocationIsEmpty tests the error when the target Location is empty.
func TestReturnErrorWhenLocationIsEmpty(t *testing.T) {
	t.Parallel()

	location, err := NewLocation(5, 5)
	assert.NoError(t, err)
	emptyLocation := Location{}

	distance, err := location.DistanceBetweenLocations(emptyLocation)
	assert.Error(t, err)
	assert.Equal(t, 0, distance)
}

// TestReturnValidRandomLocation tests the creation of a random Location instance.
func TestReturnValidRandomLocation(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		location := CreateRandomLocation()

		assert.False(t, location.IsEmpty())
		assert.GreaterOrEqual(t, location.GetX(), minCoordinate)
		assert.LessOrEqual(t, location.GetX(), maxCoordinate)
		assert.GreaterOrEqual(t, location.GetY(), minCoordinate)
		assert.LessOrEqual(t, location.GetY(), maxCoordinate)
	}
}

// TestEqualsLocation tests the equality of two Location instances.
func TestEqualsLocation(t *testing.T) {
	t.Parallel()

	location1, err := NewLocation(1, 1)
	assert.NoError(t, err)
	location2, err := NewLocation(1, 1)
	assert.NoError(t, err)

	assert.True(t, location1.Equals(location2))
}

// TestIsEmptyLocation tests the IsEmpty method of a Location instance.
func TestIsEmptyLocation(t *testing.T) {
	t.Parallel()

	var emptyLocation Location
	assert.True(t, emptyLocation.IsEmpty())
}

// TestIsSetLocation tests the IsSet method of a Location instance.
func TestIsSetLocation(t *testing.T) {
	t.Parallel()

	emptyLocation := Location{}
	assert.False(t, emptyLocation.IsSet())
}
