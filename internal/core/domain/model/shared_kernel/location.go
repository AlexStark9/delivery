package kernel

import (
	"delivery/internal/pkg/errs"
	"math/rand"
	"time"
)

const (
	minX int = 1
	minY int = 1
	maxX int = 10
	maxY int = 10
)

type Location struct {
	x int
	y int

	isSet bool
}

// NewLocation creates a new Location instance with the given x and y coordinates.
func NewLocation(x int, y int) (Location, error) {
	if x < minX || x > maxX {
		return Location{}, errs.NewValueIsOutOfRangeError(
			"The X coordinate doesn't match the boundaries: ",
			x,
			minX,
			maxX,
		)
	}
	if y < minY || y > maxY {
		return Location{}, errs.NewValueIsOutOfRangeError(
			"The X coordinate doesn't match the boundaries: ",
			y,
			minY,
			maxY,
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
	location, _ := NewLocation(minX, minY)
	return location
}

// MaxLocation creates a new Location instance with the maximum x and y coordinates.
func MaxLocation() Location {
	location, _ := NewLocation(maxX, maxY)
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

	x := rnd.Intn(maxX) + 1 // Generate a random number between 1 and maxX
	y := rnd.Intn(maxY) + 1 // Generate a random number between 1 and maxY

	return NewLocation(x, y)
}

// IsSet checks if the Location is set.
func (l Location) IsSet() bool {
	return l.isSet
}
