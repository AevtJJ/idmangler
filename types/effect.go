package types

import (
	"fmt"
)

// EffectType represents the possible types of effects
type EffectType byte

const (
	// Heal effect type
	Heal EffectType = iota
	// Mana effect type
	Mana
	// Duration effect type
	Duration
)

// String returns the string representation of an EffectType
func (e EffectType) String() string {
	switch e {
	case Heal:
		return "Heal"
	case Mana:
		return "Mana"
	case Duration:
		return "Duration"
	default:
		return fmt.Sprintf("Unknown(%d)", e)
	}
}

// BadEffectTypeError represents an error for an invalid effect type
type BadEffectTypeError struct {
	ID byte
}

// Error returns the error message for a bad effect type
func (e BadEffectTypeError) Error() string {
	return fmt.Sprintf("Invalid effect type: %d", e.ID)
}

// EffectTypeFromByte converts a byte to an EffectType or returns an error if invalid
func EffectTypeFromByte(b byte) (EffectType, error) {
	if b <= byte(Duration) {
		return EffectType(b), nil
	}
	return 0, &BadEffectTypeError{ID: b}
}

// Effect represents an effect on an item
type Effect struct {
	// Kind is the type of the effect
	Kind EffectType
	// Value is the numerical value of the effect
	Value int32
}
