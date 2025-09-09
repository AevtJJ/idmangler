package block

import (
	"fmt"

	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// DataBlockID represents the type of a data block
type DataBlockID byte

// Block ID constants
const (
	BlockStartData                 DataBlockID = 0
	BlockTypeData                  DataBlockID = 1
	BlockNameData                  DataBlockID = 2
	BlockIdentificationData        DataBlockID = 3
	BlockPowderData                DataBlockID = 4
	BlockRerollData                DataBlockID = 5
	BlockShinyData                 DataBlockID = 6
	BlockCraftedGearType           DataBlockID = 7
	BlockDurabilityData            DataBlockID = 8
	BlockRequirementsData          DataBlockID = 9
	BlockDamageData                DataBlockID = 10
	BlockDefenseData               DataBlockID = 11
	BlockCraftedIdentificationData DataBlockID = 12
	BlockCraftedConsumableTypeData DataBlockID = 13
	BlockUsesData                  DataBlockID = 14
	BlockEffectsData               DataBlockID = 15
	BlockEndData                   DataBlockID = 255
)

// String returns a string representation of the block ID
func (id DataBlockID) String() string {
	switch id {
	case BlockStartData:
		return "StartData"
	case BlockTypeData:
		return "TypeData"
	case BlockNameData:
		return "NameData"
	case BlockIdentificationData:
		return "IdentificationData"
	case BlockPowderData:
		return "PowderData"
	case BlockRerollData:
		return "RerollData"
	case BlockShinyData:
		return "ShinyData"
	case BlockCraftedGearType:
		return "CraftedGearType"
	case BlockDurabilityData:
		return "DurabilityData"
	case BlockRequirementsData:
		return "RequirementsData"
	case BlockDamageData:
		return "DamageData"
	case BlockDefenseData:
		return "DefenseData"
	case BlockCraftedIdentificationData:
		return "CraftedIdentificationData"
	case BlockCraftedConsumableTypeData:
		return "CraftedConsumableTypeData"
	case BlockUsesData:
		return "UsesData"
	case BlockEffectsData:
		return "EffectsData"
	case BlockEndData:
		return "EndData"
	default:
		return fmt.Sprintf("UnknownBlock(%d)", id)
	}
}

// InvalidBlockIDError represents an error for an invalid block ID
type InvalidBlockIDError struct {
	ID byte
}

// Error returns the error message for an invalid block ID
func (e *InvalidBlockIDError) Error() string {
	return fmt.Sprintf("Invalid block id: %d", e.ID)
}

// DataBlockIDFromByte converts a byte to a DataBlockID or returns an error if invalid
func DataBlockIDFromByte(b byte) (DataBlockID, error) {
	switch DataBlockID(b) {
	case BlockStartData, BlockTypeData, BlockNameData, BlockIdentificationData,
		BlockPowderData, BlockRerollData, BlockShinyData, BlockCraftedGearType,
		BlockDurabilityData, BlockRequirementsData, BlockDamageData, BlockDefenseData,
		BlockCraftedIdentificationData, BlockCraftedConsumableTypeData, BlockUsesData,
		BlockEffectsData, BlockEndData:
		return DataBlockID(b), nil
	default:
		return 0, &InvalidBlockIDError{ID: b}
	}
}

// DataEncoder defines the interface for encoding data blocks
type DataEncoder interface {
	// EncodeData encodes this block's data into the given output buffer
	EncodeData(ver types.EncodingVersion, out *[]byte) error

	// BlockID returns the ID of this block
	BlockID() DataBlockID
}

// DataDecoder defines the interface for decoding data blocks
type DataDecoder interface {
	// DecodeData decodes data for this block from the given bytes
	DecodeData(bytes []byte, ver types.EncodingVersion) (int, error)
}

// Block represents a data block in the encoding format
type Block interface {
	DataEncoder
	DataDecoder
}

// AnyBlock represents any type of data block
type AnyBlock interface {
	// Encode encodes this block into the given output buffer
	Encode(ver types.EncodingVersion, out *[]byte) error

	// AsID returns the ID of this block
	AsID() DataBlockID
}

// DecodeBlock decodes a single block from the given byte stream
func DecodeBlock(ver types.EncodingVersion, bytes []byte) (AnyBlock, int, error) {
	if len(bytes) == 0 {
		return nil, 0, &encoding.DecoderError{
			ErrorData: &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes},
		}
	}

	// Read the ID of the block
	id, err := DataBlockIDFromByte(bytes[0])
	if err != nil {
		return nil, 0, &encoding.DecoderError{
			ErrorData: &encoding.DecodeError{
				Type:    encoding.ErrUnknownBlock,
				Details: err,
			},
		}
	}

	// Move past the ID byte
	bytesUsed := 1
	remainingBytes := bytes[bytesUsed:]

	// Decode the block based on the ID
	var block AnyBlock
	var blockBytesUsed int

	switch id {
	case BlockStartData:
		// StartData should not be decoded after the beginning
		return nil, 0, &encoding.DecoderError{
			ErrorData: &encoding.DecodeError{Type: encoding.ErrStartReparse},
			During:    (*encoding.DataBlockID)(&id),
		}

	case BlockTypeData:
		typeData := &TypeData{}
		n, err := typeData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = typeData
		blockBytesUsed = n

	case BlockNameData:
		nameData := &NameData{}
		n, err := nameData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = nameData
		blockBytesUsed = n

	case BlockIdentificationData:
		identData := &IdentificationData{}
		n, err := identData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = identData
		blockBytesUsed = n

	case BlockPowderData:
		powderData := &PowderData{}
		n, err := powderData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = powderData
		blockBytesUsed = n

	case BlockRerollData:
		rerollData := &RerollData{}
		n, err := rerollData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = rerollData
		blockBytesUsed = n

	case BlockShinyData:
		shinyData := &ShinyData{}
		n, err := shinyData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = shinyData
		blockBytesUsed = n

	case BlockEndData:
		endData := &EndData{}
		n, err := endData.DecodeData(remainingBytes, ver)
		if err != nil {
			return nil, 0, &encoding.DecoderError{
				ErrorData: err.(*encoding.DecodeError),
				During:    (*encoding.DataBlockID)(&id),
			}
		}
		block = endData
		blockBytesUsed = n

	default:
		// Other block types not yet implemented
		return nil, 0, &encoding.DecoderError{
			ErrorData: &encoding.DecodeError{
				Type:    encoding.ErrUnknownBlock,
				Details: id,
			},
			During: (*encoding.DataBlockID)(&id),
		}
	}

	return block, bytesUsed + blockBytesUsed, nil
}

// DecodeAllBlocks decodes all blocks from the given byte stream until the end block
func DecodeAllBlocks(ver types.EncodingVersion, bytes []byte) ([]AnyBlock, error) {
	blocks := make([]AnyBlock, 0)
	bytesUsed := 0

	for bytesUsed < len(bytes) {
		block, n, err := DecodeBlock(ver, bytes[bytesUsed:])
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
		bytesUsed += n

		// If we reached the end block, stop
		if block.AsID() == BlockEndData {
			break
		}
	}

	return blocks, nil
}
