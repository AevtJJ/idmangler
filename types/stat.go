package types

// Stat represents an identification stat in Wynncraft
type Stat struct {
	// Kind is the identifier of the stat
	Kind byte
	// Base is the base value of the stat (optional when extended encoding is not used)
	Base *int32
	// Roll is the roll type of the stat (either a value or pre-identified)
	Roll RollType
}

// PreIdentified returns whether this stat is pre-identified
func (s *Stat) PreIdentified() bool {
	return s.Roll.IsPreIdentified()
}

// ContainsExtended checks if this stat contains extended data for encoding the base value
func (s *Stat) ContainsExtended() bool {
	return s.Base != nil
}

// CraftedStat represents an identification stat on a crafted item
type CraftedStat struct {
	// Kind is the identifier of the stat
	Kind byte
	// Max is the value of the stat at full durability
	Max int32
}

// NewStat creates a new Stat with the specified values
func NewStat(kind byte, base *int32, roll RollType) *Stat {
	return &Stat{
		Kind: kind,
		Base: base,
		Roll: roll,
	}
}

// NewCraftedStat creates a new CraftedStat with the specified values
func NewCraftedStat(kind byte, max int32) *CraftedStat {
	return &CraftedStat{
		Kind: kind,
		Max:  max,
	}
}
