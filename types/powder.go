package types

import (
	"fmt"
)

// Powder represents a powder with an element type and tier
type Powder struct {
	Element Element
	Tier    byte
}

// InvalidPowderTierError represents an error for an invalid powder tier
type InvalidPowderTierError struct {
	Tier byte
}

// Error returns the error message for an invalid powder tier
func (e InvalidPowderTierError) Error() string {
	return fmt.Sprintf("Invalid powder tier: %d", e.Tier)
}

// NewPowder creates a new powder with the given element and tier
// Returns an error if the tier is not between 1 and 6
func NewPowder(element Element, tier byte) (Powder, error) {
	if !ValidPowderTier(tier) {
		return Powder{}, &InvalidPowderTierError{Tier: tier}
	}
	return Powder{Element: element, Tier: tier}, nil
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
		return &InvalidPowderTierError{Tier: tier}
	}
	p.Tier = tier
	return nil
}
