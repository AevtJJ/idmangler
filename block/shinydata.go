package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// ShinyData represents a shiny stat block in an encoded item string.
// Shiny stats can be found on https://github.com/Wynntils/Static-Storage/blob/main/Data-Storage/shiny_stats.json
type ShinyData struct {
	// ID is the identifier of the shiny stat
	ID byte
	// Value is the value of the shiny stat
	Value int64
}

// BlockID returns the ID of this block
func (s *ShinyData) BlockID() DataBlockID {
	return BlockShinyData
}

// AsID returns the ID of this block
func (s *ShinyData) AsID() DataBlockID {
	return s.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (s *ShinyData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	// Write the ID byte
	*out = append(*out, s.ID)

	// Encode and write the value as a variable-length integer
	*out = append(*out, encoding.EncodeVarInt(s.Value)...)

	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (s *ShinyData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(s.BlockID()))
	// Write block data
	return s.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (s *ShinyData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	if len(bytes) < 1 {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	// Read the ID byte
	s.ID = bytes[0]
	bytesUsed := 1

	// Read the value as a variable-length integer
	var err error
	s.Value, bytesUsed, err = encoding.DecodeVarInt(bytes[1:])
	if err != nil {
		return 0, err
	}

	// Return total bytes used (ID byte + varint bytes)
	return 1 + bytesUsed, nil
}

// NewShinyData creates a new ShinyData block with the specified ID and value
func NewShinyData(id byte, value int64) *ShinyData {
	return &ShinyData{
		ID:    id,
		Value: value,
	}
}
