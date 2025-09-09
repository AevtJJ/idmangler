package types

import (
	"fmt"
)

// ItemType represents the different types of items
type ItemType byte

const (
	// Gear represents equipment items
	Gear ItemType = iota
	// Tome represents tome items
	Tome
	// Charm represents charm items
	Charm
	// CraftedGear represents crafted equipment items
	CraftedGear
	// CraftedConsu represents crafted consumable items
	CraftedConsu
)

// String returns the string representation of an ItemType
func (i ItemType) String() string {
	switch i {
	case Gear:
		return "Gear"
	case Tome:
		return "Tome"
	case Charm:
		return "Charm"
	case CraftedGear:
		return "CraftedGear"
	case CraftedConsu:
		return "CraftedConsu"
	default:
		return fmt.Sprintf("Unknown(%d)", i)
	}
}

// BadItemTypeError represents an error for an invalid item type ID
type BadItemTypeError struct {
	ID byte
}

// Error returns the error message for a bad item type
func (e BadItemTypeError) Error() string {
	return fmt.Sprintf("Invalid item type id: %d", e.ID)
}

// ItemTypeFromByte converts a byte to an ItemType or returns an error if invalid
func ItemTypeFromByte(b byte) (ItemType, error) {
	if b <= byte(CraftedConsu) {
		return ItemType(b), nil
	}
	return 0, &BadItemTypeError{ID: b}
}
