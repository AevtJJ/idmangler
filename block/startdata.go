package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// StartData represents the start block of an encoded item string.
// It contains information about the encoding version to be used.
type StartData struct {
	Version types.EncodingVersion
}

// BlockID returns the ID of this block
func (s *StartData) BlockID() DataBlockID {
	return BlockStartData
}

// AsID returns the ID of this block
func (s *StartData) AsID() DataBlockID {
	return s.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (s *StartData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	*out = append(*out, byte(s.Version))
	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (s *StartData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(s.BlockID()))
	// Write block data
	return s.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (s *StartData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	// StartData cannot be decoded after the start of the encoding
	// as it should only appear at the beginning
	return 0, &encoding.DecodeError{Type: encoding.ErrStartReparse}
}

// DecodeStartBytes is a special case function for parsing the start bytes
func DecodeStartBytes(bytes []byte) (*StartData, int, error) {
	if len(bytes) < 2 {
		return nil, 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	idByte := bytes[0]
	if idByte != byte(BlockStartData) {
		return nil, 0, &encoding.DecodeError{Type: encoding.ErrNoStartBlockFound}
	}

	verByte := bytes[1]
	ver, err := types.EncodingVersionFromByte(verByte)
	if err != nil {
		return nil, 0, &encoding.DecodeError{
			Type:    encoding.ErrUnknownVersion,
			Details: verByte,
		}
	}

	return &StartData{Version: ver}, 2, nil
}

// NewStartData creates a new StartData block with the specified version
func NewStartData(version types.EncodingVersion) *StartData {
	return &StartData{
		Version: version,
	}
}
