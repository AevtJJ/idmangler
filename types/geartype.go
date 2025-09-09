package types

import (
	"fmt"
)

// CraftedGearType represents the possible types of gear items
type CraftedGearType byte

const (
	// Spear gear type
	Spear CraftedGearType = iota
	// Wand gear type
	Wand
	// Dagger gear type
	Dagger
	// Bow gear type
	Bow
	// Relik gear type
	Relik

	// Ring gear type
	Ring
	// Bracelet gear type
	Bracelet
	// Necklace gear type
	Necklace

	// Helmet gear type
	Helmet
	// Chestplate gear type
	Chestplate
	// Leggings gear type
	Leggings
	// Boots gear type
	Boots

	// Weapon is a fallback for "signed, crafted gear with a skin"
	Weapon = 12
	// Accessory is a fallback for when specific gear type is not known
	Accessory = 13
)

// String returns the string representation of a CraftedGearType
func (g CraftedGearType) String() string {
	switch g {
	case Spear:
		return "Spear"
	case Wand:
		return "Wand"
	case Dagger:
		return "Dagger"
	case Bow:
		return "Bow"
	case Relik:
		return "Relik"
	case Ring:
		return "Ring"
	case Bracelet:
		return "Bracelet"
	case Necklace:
		return "Necklace"
	case Helmet:
		return "Helmet"
	case Chestplate:
		return "Chestplate"
	case Leggings:
		return "Leggings"
	case Boots:
		return "Boots"
	case Weapon:
		return "Weapon"
	case Accessory:
		return "Accessory"
	default:
		return fmt.Sprintf("Unknown(%d)", g)
	}
}

// BadGearTypeError represents an error for an invalid gear type ID
type BadGearTypeError struct {
	ID byte
}

// Error returns the error message for a bad gear type
func (e BadGearTypeError) Error() string {
	return fmt.Sprintf("Invalid gear type id: %d", e.ID)
}

// CraftedGearTypeFromByte converts a byte to a CraftedGearType or returns an error if invalid
func CraftedGearTypeFromByte(b byte) (CraftedGearType, error) {
	if b <= 13 && (b <= 11 || b >= 12) { // Valid values: 0-11, 12, 13
		return CraftedGearType(b), nil
	}
	return 0, &BadGearTypeError{ID: b}
}
