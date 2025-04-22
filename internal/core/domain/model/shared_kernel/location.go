package shared_kernel

import (
	"delivery/internal/pkg/errs"
	"math"
	"math/rand"
	"time"
)

// Constant min and max value for location
const (
	minCoordinate int = 1
	maxCoordinate int = 10
)

// Location represents a Value Object with X and Y coordinates
type Location struct {
	x int
	y int

	isSet bool
}

// NewLocation creates a new Location instance with the given x and y coordinates.
func NewLocation(x int, y int) (Location, error) {
	if x < minCoordinate || x > maxCoordinate {
		return Location{}, errs.NewValueIsOutOfRangeError(
			"The X coordinate doesn't match the boundaries: ",
			x,
			minCoordinate,
			maxCoordinate,
		)
	}
	if y < minCoordinate || y > maxCoordinate {
		return Location{}, errs.NewValueIsOutOfRangeError(
			"The X coordinate doesn't match the boundaries: ",
			y,
			minCoordinate,
			maxCoordinate,
		)
	}

	return Location{
		x: x,
		y: y,

		isSet: true,
	}, nil
}

// MinLocation creates a new Location instance with the minimum x and y coordinates.
func MinLocation() Location {
	location, _ := NewLocation(minCoordinate, minCoordinate)
	return location
}

// MaxLocation creates a new Location instance with the maximum x and y coordinates.
func MaxLocation() Location {
	location, _ := NewLocation(maxCoordinate, maxCoordinate)
	return location
}

// GetX returns the x coordinate of the Location.
func (l Location) GetX() int {
	return l.x
}

// GetY returns the y coordinate of the Location.
func (l Location) GetY() int {
	return l.y
}

// CreateRandomLocation creates a new Location instance with random x and y coordinates.
func CreateRandomLocation() (Location, error) {
	// Create a new random source
	scr := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(scr)
	// Generate a random number between 1 and maxCoordinate
	x := rnd.Intn(maxCoordinate) + 1
	y := rnd.Intn(maxCoordinate) + 1

	return NewLocation(x, y)
}

// EqualsLocation checks if two Locations are equivalent.
func (l Location) EqualsLocation(otherLocation Location) bool {
	return l == otherLocation
}

// IsEmpty checks if the Location is empty.
func (l Location) IsEmpty() bool {
	return !l.isSet
}

// IsSet checks if the Location is set.
func (l Location) IsSet() bool {
	return l.isSet
}

// DistanceBetweenLocations calculates the distance between two Locations.
func (l Location) DistanceBetweenLocations(otherLocation Location) (int, error) {
	if otherLocation.IsEmpty() {
		return 0, errs.NewValueIsRequiredError("otherLocation")
	}
	return int(math.Abs(float64(l.x-otherLocation.x)) + math.Abs(float64(l.y-otherLocation.y))), nil
}
