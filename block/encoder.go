package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// ItemEncoder handles the encoding of blocks to ID strings
type ItemEncoder struct {
	version types.EncodingVersion
}

// NewItemEncoder creates a new ItemEncoder with the default version
func NewItemEncoder() *ItemEncoder {
	return &ItemEncoder{
		version: types.Version1,
	}
}

// WithVersion sets the encoding version and returns the encoder
func (e *ItemEncoder) WithVersion(version types.EncodingVersion) *ItemEncoder {
	e.version = version
	return e
}

// EncodeBlocks encodes a series of blocks to an ID string
func (e *ItemEncoder) EncodeBlocks(blocks []AnyBlock) (string, error) {
	// Ensure we have at least a start block
	if len(blocks) == 0 || blocks[0].AsID() != BlockStartData {
		startBlock := NewStartData(e.version)
		blocks = append([]AnyBlock{startBlock}, blocks...)
	} else {
		// Update the version in the start block
		if startBlock, ok := blocks[0].(*StartData); ok {
			e.version = startBlock.Version
		}
	}

	// Ensure we have an end block at the end
	if len(blocks) == 0 || blocks[len(blocks)-1].AsID() != BlockEndData {
		endBlock := NewEndData()
		blocks = append(blocks, endBlock)
	}

	// Encode all blocks
	var bytes []byte
	for _, b := range blocks {
		if err := b.Encode(e.version, &bytes); err != nil {
			return "", err
		}
	}

	// Convert bytes to string
	return encoding.EncodeString(bytes), nil
}

// ItemDecoder handles the decoding of ID strings to blocks
type ItemDecoder struct{}

// NewItemDecoder creates a new ItemDecoder
func NewItemDecoder() *ItemDecoder {
	return &ItemDecoder{}
}

// DecodeString decodes an ID string into a series of blocks
func (d *ItemDecoder) DecodeString(idString string) ([]AnyBlock, error) {
	// Convert string to bytes
	bytes, err := encoding.DecodeString(idString)
	if err != nil {
		return nil, err
	}

	// Start by decoding the start block to get the version
	startBlock, bytesRead, err := DecodeStartBytes(bytes)
	if err != nil {
		return nil, err
	}

	// Decode the remaining blocks
	remainingBlocks, err := DecodeAllBlocks(startBlock.Version, bytes[bytesRead:])
	if err != nil {
		return nil, err
	}

	// Combine start block with remaining blocks
	allBlocks := make([]AnyBlock, 0, len(remainingBlocks)+1)
	allBlocks = append(allBlocks, startBlock)
	allBlocks = append(allBlocks, remainingBlocks...)

	return allBlocks, nil
}
