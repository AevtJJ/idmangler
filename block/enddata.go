package block

import (
	"github.com/AevtJJ/idmangler/types"
)

// EndData represents the end block of an encoded item string.
// This block is always empty and marks the end of the encoded data.
type EndData struct{}

// BlockID returns the ID of this block
func (e *EndData) BlockID() DataBlockID {
	return BlockEndData
}

// AsID returns the ID of this block
func (e *EndData) AsID() DataBlockID {
	return e.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
// EndData is always empty, so this does nothing
func (e *EndData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	// EndData is always empty
	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (e *EndData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(e.BlockID()))
	// Write block data (which is empty)
	return e.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
// EndData is always empty, so this just returns successfully
func (e *EndData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	// EndData is always empty
	return 0, nil
}

// NewEndData creates a new EndData block
func NewEndData() *EndData {
	return &EndData{}
}
