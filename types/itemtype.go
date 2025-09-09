package types

import (
	"fmt"
	"strings"
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

// ItemTypeFromString converts a string to an ItemType
func ItemTypeFromString(s string) ItemType {
	switch strings.ToLower(s) {
	case "gear":
		return Gear
	case "tome":
		return Tome
	case "charm":
		return Charm
	case "craftedgear":
		return CraftedGear
	case "craftedconsu":
		return CraftedConsu
	default:
		return Gear // Default to Gear if unknown
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
