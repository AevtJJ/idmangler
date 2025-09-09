package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// NameData represents the name block of an encoded item string.
// It stores the item name as an ASCII string.
type NameData struct {
	Name string
}

// BlockID returns the ID of this block
func (n *NameData) BlockID() DataBlockID {
	return BlockNameData
}

// AsID returns the ID of this block
func (n *NameData) AsID() DataBlockID {
	return n.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (n *NameData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	// Check that the string is valid ASCII
	for _, c := range n.Name {
		if c > 127 {
			return &encoding.EncodeError{Type: encoding.ErrNonAsciiString}
		}
	}

	// Append the string bytes
	*out = append(*out, []byte(n.Name)...)
	// Append the null terminator
	*out = append(*out, 0)
	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (n *NameData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(n.BlockID()))
	// Write block data
	return n.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (n *NameData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	// Find the null terminator
	nullPos := -1
	for i, b := range bytes {
		if b == 0 {
			nullPos = i
			break
		}
	}

	if nullPos == -1 {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	// Extract the name (without the null terminator)
	n.Name = string(bytes[:nullPos])

	// Return bytes used (including the null terminator)
	return nullPos + 1, nil
}

// NewNameData creates a new NameData block with the specified name
func NewNameData(name string) *NameData {
	return &NameData{Name: name}
}
