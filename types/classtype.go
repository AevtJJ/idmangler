package types

import (
	"fmt"
)

// ClassType represents the character classes
type ClassType byte

const (
	// ClassUnspecified represents no class specified (0)
	ClassUnspecified ClassType = iota
	// Mage class
	Mage
	// Archer class
	Archer
	// Warrior class
	Warrior
	// Assasin class
	Assasin
	// Shaman class
	Shaman
)

// String returns the string representation of a ClassType
func (c ClassType) String() string {
	switch c {
	case Mage:
		return "Mage"
	case Archer:
		return "Archer"
	case Warrior:
		return "Warrior"
	case Assasin:
		return "Assasin"
	case Shaman:
		return "Shaman"
	case ClassUnspecified:
		return "Unspecified"
	default:
		return fmt.Sprintf("Unknown(%d)", c)
	}
}

// BadClassTypeError represents an error for an invalid class type ID
type BadClassTypeError struct {
	ID byte
}

// Error returns the error message for a bad class type
func (e BadClassTypeError) Error() string {
	return fmt.Sprintf("Invalid class type id: %d", e.ID)
}

// ClassTypeFromByte converts a byte to a ClassType or returns an error if invalid
func ClassTypeFromByte(b byte) (ClassType, error) {
	if b <= byte(Shaman) {
		return ClassType(b), nil
	}
	return 0, &BadClassTypeError{ID: b}
}
