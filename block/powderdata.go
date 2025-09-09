package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// PowderData represents powder data for an item in an encoded item string.
type PowderData struct {
	// PowderSlots is the number of powder slots available on this item
	PowderSlots byte
	// Powders is a slice of powders applied to this item
	Powders []types.Powder
}

// BlockID returns the ID of this block
func (p *PowderData) BlockID() DataBlockID {
	return BlockPowderData
}

// AsID returns the ID of this block
func (p *PowderData) AsID() DataBlockID {
	return p.BlockID()
}

// EncodeData encodes this block's data into the given output buffer
func (p *PowderData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	// Check if we have too many powders
	if len(p.Powders) > 255 {
		return &encoding.EncodeError{Type: encoding.ErrTooManyPowders}
	}

	// Calculate how many bytes are needed for the powders
	// Each powder uses 5 bits (3 for element, 2 for tier)
	bitsNeeded := len(p.Powders) * 5
	totalBytes := (bitsNeeded + 7) / 8

	// Create the powder data bytes
	powderData := make([]byte, totalBytes)

	// Encode each powder
	for i, powder := range p.Powders {
		// Calculate the 5-bit powder value: (element * 6 + tier) & 0b00011111
		elem := byte(powder.Element)
		tier := powder.Tier
		powderNum := (elem*6 + tier) & 0b00011111

		// Bit position where this specific powder starts
		powderIdx := i * 5

		// Set the bits for this powder
		for j := 0; j < 5; j++ {
			// Calculate the bit position of this bit
			idx := powderIdx + j
			bit := (powderNum >> (4 - j)) & 0b1
			powderData[idx/8] |= bit << (7 - (idx % 8))
		}
	}

	// Write the powder slots
	*out = append(*out, p.PowderSlots)
	// Write the number of powders
	*out = append(*out, byte(len(p.Powders)))
	// Write the powder data
	*out = append(*out, powderData...)

	return nil
}

// Encode encodes this block with its ID into the given output buffer
func (p *PowderData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write block ID
	*out = append(*out, byte(p.BlockID()))
	// Write block data
	return p.EncodeData(ver, out)
}

// DecodeData decodes data for this block from the given bytes
func (p *PowderData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	if len(bytes) < 2 {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	// Read the powder slots
	p.PowderSlots = bytes[0]

	// Read the number of powders
	powderCount := int(bytes[1])
	bytesUsed := 2

	// Calculate how many bytes we need to read for the powders
	bitsNeeded := powderCount * 5
	totalBytes := (bitsNeeded + 7) / 8

	if len(bytes) < bytesUsed+totalBytes {
		return 0, &encoding.DecodeError{Type: encoding.ErrUnexpectedEndOfBytes}
	}

	// Read the powder data bytes
	powderDataBytes := bytes[bytesUsed : bytesUsed+totalBytes]
	bytesUsed += totalBytes

	// Initialize the powders slice
	p.Powders = make([]types.Powder, 0, powderCount)

	// Decode each powder
	for powderIdx := 0; powderIdx < powderCount; powderIdx++ {
		var powderValue byte

		// Extract the 5 bits for this powder
		for i := 0; i < 5; i++ {
			idx := (powderIdx * 5) + i
			bit := (powderDataBytes[idx/8] >> (7 - (idx % 8))) & 0b1
			powderValue |= bit << (4 - i)
		}

		// Skip empty powders
		if powderValue == 0 {
			continue
		}

		// Calculate element and tier from the powder value
		var elem, tier byte
		if powderValue%6 == 0 {
			elem = (powderValue / 6) - 1
			tier = 6
		} else {
			elem = powderValue / 6
			tier = powderValue % 6
		}

		// Create the powder
		element, err := types.ElementFromByte(elem)
		if err != nil {
			return 0, &encoding.DecodeError{
				Type:    encoding.ErrBadElement,
				Details: elem,
			}
		}

		powder, err := types.NewPowder(element, tier)
		if err != nil {
			return 0, &encoding.DecodeError{
				Type:    encoding.ErrBadPowderTier,
				Details: tier,
			}
		}

		p.Powders = append(p.Powders, powder)
	}

	return bytesUsed, nil
}

// NewPowderData creates a new PowderData block with the specified powder slots and powders
func NewPowderData(powderSlots byte, powders []types.Powder) *PowderData {
	return &PowderData{
		PowderSlots: powderSlots,
		Powders:     powders,
	}
}
