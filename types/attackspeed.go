package types

import (
	"fmt"
)

// AttackSpeed represents the attack speed of a weapon
type AttackSpeed byte

const (
	// SuperFast attack speed
	SuperFast AttackSpeed = iota
	// VeryFast attack speed
	VeryFast
	// Fast attack speed
	Fast
	// Normal attack speed
	Normal
	// Slow attack speed
	Slow
	// VerySlow attack speed
	VerySlow
	// SuperSlow attack speed
	SuperSlow
)

// String returns the string representation of an AttackSpeed
func (a AttackSpeed) String() string {
	switch a {
	case SuperFast:
		return "SuperFast"
	case VeryFast:
		return "VeryFast"
	case Fast:
		return "Fast"
	case Normal:
		return "Normal"
	case Slow:
		return "Slow"
	case VerySlow:
		return "VerySlow"
	case SuperSlow:
		return "SuperSlow"
	default:
		return fmt.Sprintf("Unknown(%d)", a)
	}
}

// BadAttackSpeedError represents an error for an invalid attack speed ID
type BadAttackSpeedError struct {
	ID byte
}

// Error returns the error message for a bad attack speed
func (e BadAttackSpeedError) Error() string {
	return fmt.Sprintf("Invalid attack speed id: %d", e.ID)
}

// AttackSpeedFromByte converts a byte to an AttackSpeed or returns an error if invalid
func AttackSpeedFromByte(b byte) (AttackSpeed, error) {
	if b <= byte(SuperSlow) {
		return AttackSpeed(b), nil
	}
	return 0, &BadAttackSpeedError{ID: b}
}

// Compare compares this attack speed with another
// Returns:
//
//	-1 if this is faster than other
//	 0 if they are equal
//	 1 if this is slower than other
func (a AttackSpeed) Compare(other AttackSpeed) int {
	// Attack speed ordering is reversed (lower values are faster)
	if a < other {
		return -1
	} else if a > other {
		return 1
	}
	return 0
}
