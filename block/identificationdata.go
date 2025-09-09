package block

import (
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// IdentificationData represents a block containing item identification data
type IdentificationData struct {
	// The identifications
	Identifications []*types.Stat
	// Whether or not extended encoding is used or to be used for encoding.
	// If extended encoding is used then all values will have their base values and rolls encoded.
	// Without extended encoding only the rolls are encoded and pre-identified values are ignored.
	ExtendedEncoding bool
}

// BlockID returns the ID of this block
func (d *IdentificationData) BlockID() DataBlockID {
	return BlockIdentificationData
}

// AsID returns the ID of this block
func (d *IdentificationData) AsID() DataBlockID {
	return d.BlockID()
}

// Encode encodes this block into the given output buffer
func (d *IdentificationData) Encode(ver types.EncodingVersion, out *[]byte) error {
	// Write the block ID
	*out = append(*out, byte(d.BlockID()))

	// Encode the data
	return d.EncodeData(ver, out)
}

// EncodeData encodes this block's data into the given output buffer
func (d *IdentificationData) EncodeData(ver types.EncodingVersion, out *[]byte) error {
	switch ver {
	case types.Version1:
		// Count non pre-identified stats
		nonPreIdCount := 0
		preIdCount := 0
		for _, id := range d.Identifications {
			if !id.PreIdentified() {
				nonPreIdCount++
			} else {
				preIdCount++
			}
		}

		// Check for too many identifications
		if nonPreIdCount > 255 || preIdCount > 255 {
			return &encoding.EncodeError{
				Type: encoding.ErrTooManyIdentifications,
			}
		}

		// Add number of non pre-identified stats
		*out = append(*out, byte(nonPreIdCount))

		// Add extended encoding flag
		if d.ExtendedEncoding {
			*out = append(*out, 1)
		} else {
			*out = append(*out, 0)
		}

		return d.encodeIndividualIdents(out)
	default:
		return &encoding.EncodeError{
			Type:    encoding.ErrUnknownVersion,
			Details: ver,
		}
	}
}

// encodeIndividualIdents encodes the individual identification stats
func (d *IdentificationData) encodeIndividualIdents(out *[]byte) error {
	// Encode pre-identified stats if using extended encoding
	if d.ExtendedEncoding {
		// Count pre-identified stats
		preIdStats := make([]*types.Stat, 0)
		for _, id := range d.Identifications {
			if id.PreIdentified() {
				preIdStats = append(preIdStats, id)
			}
		}

		// Add count of pre-identified stats
		*out = append(*out, byte(len(preIdStats)))

		// Add pre-identified stats
		for _, stat := range preIdStats {
			// Add the ID of the stat
			*out = append(*out, stat.Kind)

			// Add the base value
			if stat.Base == nil {
				return &encoding.EncodeError{
					Type:    encoding.ErrNoBasevalueGiven,
					Details: stat.Kind,
				}
			}
			*out = append(*out, encoding.EncodeVarInt(int64(*stat.Base))...)
		}
	}

	// Encode non-preidentified stats
	for _, ident := range d.Identifications {
		// Only handle non pre-identified stats
		if !ident.PreIdentified() {
			// Add ID of the stat
			*out = append(*out, ident.Kind)

			// Add base value if using extended encoding
			if d.ExtendedEncoding {
				if ident.Base == nil {
					return &encoding.EncodeError{
						Type:    encoding.ErrNoBasevalueGiven,
						Details: ident.Kind,
					}
				}
				*out = append(*out, encoding.EncodeVarInt(int64(*ident.Base))...)
			}

			// Add roll value
			*out = append(*out, ident.Roll.Value())
		}
	}

	return nil
}

// DecodeData decodes data for this block from the given bytes
func (d *IdentificationData) DecodeData(bytes []byte, ver types.EncodingVersion) (int, error) {
	if len(bytes) < 2 {
		return 0, &encoding.DecodeError{
			Type: encoding.ErrUnexpectedEndOfBytes,
		}
	}

	switch ver {
	case types.Version1:
		var bytesUsed int

		// First byte is the number of identifications
		identCount := int(bytes[0])

		// Second byte is whether extended encoding is used
		extendedEncoding := bytes[1] == 1

		bytesUsed = 2

		// Create slice for identifications
		idents := make([]*types.Stat, 0)

		var preIdCount int
		if extendedEncoding {
			// If extended encoding, next byte is the count of pre-identified stats
			if len(bytes) <= bytesUsed {
				return bytesUsed, &encoding.DecodeError{
					Type: encoding.ErrUnexpectedEndOfBytes,
				}
			}
			preIdCount = int(bytes[bytesUsed])
			bytesUsed++

			// Decode pre-identified stats
			for i := 0; i < preIdCount; i++ {
				// Get the stat ID
				if len(bytes) <= bytesUsed {
					return bytesUsed, &encoding.DecodeError{
						Type: encoding.ErrUnexpectedEndOfBytes,
					}
				}
				id := bytes[bytesUsed]
				bytesUsed++

				// Decode the base value
				baseVal, n, err := encoding.DecodeVarInt(bytes[bytesUsed:])
				if err != nil {
					return bytesUsed, err
				}
				bytesUsed += n

				// Create the stat
				base := int32(baseVal)
				idents = append(idents, types.NewStat(id, &base, types.NewPreIdentifiedRoll()))
			}
		}

		// Decode non-preidentified stats
		for i := 0; i < identCount; i++ {
			// Get the stat ID
			if len(bytes) <= bytesUsed {
				return bytesUsed, &encoding.DecodeError{
					Type: encoding.ErrUnexpectedEndOfBytes,
				}
			}
			id := bytes[bytesUsed]
			bytesUsed++

			// Decode the base value if extended encoding is used
			var baseVal *int32
			if extendedEncoding {
				val, n, err := encoding.DecodeVarInt(bytes[bytesUsed:])
				if err != nil {
					return bytesUsed, err
				}
				bytesUsed += n
				intVal := int32(val)
				baseVal = &intVal
			}

			// Get the roll value
			if len(bytes) <= bytesUsed {
				return bytesUsed, &encoding.DecodeError{
					Type: encoding.ErrUnexpectedEndOfBytes,
				}
			}
			rollVal := bytes[bytesUsed]
			bytesUsed++

			// Create the stat
			idents = append(idents, types.NewStat(id, baseVal, types.NewValueRoll(rollVal)))
		}

		// Store the decoded data
		d.Identifications = idents
		d.ExtendedEncoding = extendedEncoding

		return bytesUsed, nil

	default:
		return 0, &encoding.DecodeError{
			Type:    encoding.ErrUnknownVersion,
			Details: ver,
		}
	}
}

// NewIdentificationData creates a new IdentificationData block with the given stats and encoding mode
func NewIdentificationData(identifications []*types.Stat, extendedEncoding bool) *IdentificationData {
	return &IdentificationData{
		Identifications:  identifications,
		ExtendedEncoding: extendedEncoding,
	}
}
