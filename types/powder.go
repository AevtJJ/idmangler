package types

import (
	"fmt"
)

// Powder represents a powder that can be applied to an item
type Powder struct {
	// Element is the elemental type of the powder
	Element Element
	// Tier is the tier of the powder (1-6)
	Tier byte
}

// String returns a string representation of the powder
func (p Powder) String() string {
	return fmt.Sprintf("%s T%d", p.Element, p.Tier)
}

// NewPowder creates a new powder with the specified element and tier
func NewPowder(element Element, tier byte) (Powder, error) {
	if !ValidPowderTier(tier) {
		return Powder{}, &BadPowderTierError{Tier: tier}
	}
	return Powder{
		Element: element,
		Tier:    tier,
	}, nil
}

// BadPowderTierError represents an error for an invalid powder tier
type BadPowderTierError struct {
	Tier byte
}

// Error returns the error message for a bad powder tier
func (e BadPowderTierError) Error() string {
	return fmt.Sprintf("Invalid powder tier: %d", e.Tier)
}

// ValidPowderTier checks if the given tier is valid for a powder (1-6)
func ValidPowderTier(tier byte) bool {
	return tier >= 1 && tier <= 6
}

// SetElement sets the element type of the powder
func (p *Powder) SetElement(element Element) {
	p.Element = element
}

// SetTier sets the tier of the powder
// Returns an error if the tier is not between 1 and 6
func (p *Powder) SetTier(tier byte) error {
	if !ValidPowderTier(tier) {
		return &BadPowderTierError{Tier: tier}
	}
	p.Tier = tier
	return nil
}
