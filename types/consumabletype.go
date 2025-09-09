package types

import (
	"fmt"
)

// ConsumableType represents the possible types of consumables
type ConsumableType byte

const (
	// Potion consumable type
	Potion ConsumableType = iota
	// Food consumable type
	Food
	// Scroll consumable type
	Scroll
)

// String returns the string representation of a ConsumableType
func (c ConsumableType) String() string {
	switch c {
	case Potion:
		return "Potion"
	case Food:
		return "Food"
	case Scroll:
		return "Scroll"
	default:
		return fmt.Sprintf("Unknown(%d)", c)
	}
}

// BadConsumableTypeError represents an error for an invalid consumable type ID
type BadConsumableTypeError struct {
	ID byte
}

// Error returns the error message for a bad consumable type
func (e BadConsumableTypeError) Error() string {
	return fmt.Sprintf("Invalid consumable type id: %d", e.ID)
}

// ConsumableTypeFromByte converts a byte to a ConsumableType or returns an error if invalid
func ConsumableTypeFromByte(b byte) (ConsumableType, error) {
	if b <= byte(Scroll) {
		return ConsumableType(b), nil
	}
	return 0, &BadConsumableTypeError{ID: b}
}
