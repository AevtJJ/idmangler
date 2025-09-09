package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// TypeData represents the item type block of an encoded item string.
type TypeData struct {
	ItemType types.ItemType
}

// BlockID returns the ID of this block
func (t *TypeData) BlockID() DataBlockID {
	return BlockTypeData
}

// AsID returns the ID of this block
func (t *TypeData) AsID() DataBlockID {
	return t.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (t *TypeData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	*out = append(*out, byte(t.ItemType))
	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (t *TypeData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(t.BlockID()))
	// Write block data
	return t.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (t *TypeData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	if len(bytes) < 1 {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	itemType, err := types.ItemTypeFromByte(bytes[0])
	if err != nil {
		return 0, &encoding.DecodeError{
			Type:    encoding.ErrBadItemType,
			Details: bytes[0],
		}
	}

	t.ItemType = itemType
	return 1, nil
}

// NewTypeData creates a new TypeData block with the specified item type
func NewTypeData(itemType types.ItemType) *TypeData {
	return &TypeData{
		ItemType: itemType,
	}
}
