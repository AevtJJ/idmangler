package types

import (
	"fmt"
)

// Element represents the elemental types
type Element byte

const (
	// Earth element
	Earth Element = iota
	// Thunder element
	Thunder
	// Water element
	Water
	// Fire element
	Fire
	// Air element
	Air
)

// String returns the string representation of an Element
func (e Element) String() string {
	switch e {
	case Earth:
		return "Earth"
	case Thunder:
		return "Thunder"
	case Water:
		return "Water"
	case Fire:
		return "Fire"
	case Air:
		return "Air"
	default:
		return fmt.Sprintf("Unknown(%d)", e)
	}
}

// BadElementError represents an error for an invalid element ID
type BadElementError struct {
	ID byte
}

// Error returns the error message for a bad element
func (e BadElementError) Error() string {
	return fmt.Sprintf("Invalid element id: %d", e.ID)
}

// ElementFromByte converts a byte to an Element or returns an error if invalid
func ElementFromByte(b byte) (Element, error) {
	if b <= byte(Air) {
		return Element(b), nil
	}
	return 0, &BadElementError{ID: b}
}
