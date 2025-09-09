package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// RerollData represents a reroll data block in an encoded item string.
// It stores information about how many times an item has been rerolled.
type RerollData struct {
	// Rerolls is the number of times the item has been rerolled
	Rerolls byte
}

// BlockID returns the ID of this block
func (r *RerollData) BlockID() DataBlockID {
	return BlockRerollData
}

// AsID returns the ID of this block
func (r *RerollData) AsID() DataBlockID {
	return r.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (r *RerollData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	// Write the reroll count
	*out = append(*out, r.Rerolls)
	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (r *RerollData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(r.BlockID()))
	// Write block data
	return r.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (r *RerollData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	if len(bytes) < 1 {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	// Read the reroll count
	r.Rerolls = bytes[0]
	return 1, nil
}

// NewRerollData creates a new RerollData block with the specified reroll count
func NewRerollData(rerolls byte) *RerollData {
	return &RerollData{
		Rerolls: rerolls,
	}
}
