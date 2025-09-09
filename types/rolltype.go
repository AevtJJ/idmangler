package types

// RollType represents the possible roll types of a Stat
type RollType interface {
	// IsPreIdentified returns whether this roll type is pre-identified
	IsPreIdentified() bool
	// Value returns the roll value if available, 0 otherwise
	Value() byte
}

// ValueRoll represents a roll between 0% and 255% of the base value
type ValueRoll struct {
	// Val is the roll value (0-255)
	Val byte
}

// IsPreIdentified returns false for a ValueRoll
func (v *ValueRoll) IsPreIdentified() bool {
	return false
}

// Value returns the roll value
func (v *ValueRoll) Value() byte {
	return v.Val
}

// PreIdentifiedRoll represents a fixed roll value
type PreIdentifiedRoll struct{}

// IsPreIdentified returns true for a PreIdentifiedRoll
func (p *PreIdentifiedRoll) IsPreIdentified() bool {
	return true
}

// Value returns 0 for a PreIdentifiedRoll
func (p *PreIdentifiedRoll) Value() byte {
	return 0
}

// NewValueRoll creates a new ValueRoll with the specified value
func NewValueRoll(value byte) *ValueRoll {
	return &ValueRoll{Val: value}
}

// NewPreIdentifiedRoll creates a new PreIdentifiedRoll
func NewPreIdentifiedRoll() *PreIdentifiedRoll {
	return &PreIdentifiedRoll{}
}
