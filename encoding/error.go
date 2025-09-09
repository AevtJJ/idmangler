package encoding

import (
	"fmt"
)

// DataBlockID represents the type of data block (will be fully defined in the block package)
type DataBlockID byte

// EncodeError represents errors that can occur during encoding
type EncodeError struct {
	Type    EncodeErrorType
	Details string
}

// EncodeErrorType identifies the type of encoding error that occurred
type EncodeErrorType int

const (
	// ErrNoStartBlock indicates no start data block was found when encoding
	ErrNoStartBlock EncodeErrorType = iota
	// ErrNonAsciiString indicates encoder was given a string with non-ASCII characters
	ErrNonAsciiString
	// ErrTooManyIdentifications indicates more than 255 identifications were passed for encoding
	ErrTooManyIdentifications
	// ErrNoBaseValueGiven indicates identification is missing a base value while using extended encoding
	ErrNoBaseValueGiven
	// ErrTooManyPowders indicates more than 255 powders were passed for encoding
	ErrTooManyPowders
	// ErrEffectStrengthTooHigh indicates effect strength should be a percentage between 0 and 100
	ErrEffectStrengthTooHigh
	// ErrTooManySkills indicates more than 255 skills were passed for encoding
	ErrTooManySkills
	// ErrTooManyDamageValues indicates more than 255 damage values were passed for encoding
	ErrTooManyDamageValues
	// ErrTooManyEffects indicates more than 255 effects were passed for encoding
	ErrTooManyEffects
	// ErrTooManyDefenses indicates more than 255 defense values were passed for encoding
	ErrTooManyDefenses
)

// Error returns the error message for an EncodeError
func (e *EncodeError) Error() string {
	switch e.Type {
	case ErrNoStartBlock:
		return "No start data block found"
	case ErrNonAsciiString:
		return "Cannot encode non ASCII string"
	case ErrTooManyIdentifications:
		return "Cannot encode more than 255 identifications per item"
	case ErrNoBaseValueGiven:
		return fmt.Sprintf("Identification id: %s was not given a base value while using extended encoding", e.Details)
	case ErrTooManyPowders:
		return "Cannot encode more than 255 powders per item"
	case ErrEffectStrengthTooHigh:
		return fmt.Sprintf("Effect strength of %s is too high, it should be a percentage between 0 and 100", e.Details)
	case ErrTooManySkills:
		return "Cannot encode more than 255 skills per item"
	case ErrTooManyDamageValues:
		return "Cannot encode more than 255 damage values per item"
	case ErrTooManyEffects:
		return "Cannot encode more than 255 effects per item"
	case ErrTooManyDefenses:
		return "Cannot encode more than 255 defenses per item"
	default:
		return "Unknown encoding error"
	}
}

// EncoderError provides context about which block caused an encoding error
type EncoderError struct {
	ErrorData *EncodeError
	During    DataBlockID
}

// Error returns the error message for an EncoderError
func (e *EncoderError) Error() string {
	return fmt.Sprintf("%s While encoding block %d", e.Error(), e.During)
}

// DecodeErrorType identifies the type of decoding error
type DecodeErrorType int

const (
	// ErrNoStartBlockFound indicates the ID string doesn't start with a valid start block
	ErrNoStartBlockFound DecodeErrorType = iota
	// ErrUnknownVersion indicates encoding of an unknown potentially future version was hit
	ErrUnknownVersion
	// ErrStartReparse indicates decoder found a second start block in the data
	ErrStartReparse
	// ErrUnknownBlock indicates decoder hit an unknown block it could not decode
	ErrUnknownBlock
	// ErrBadString indicates an invalid non-ascii/utf-8 string was decoded
	ErrBadString
	// ErrBadItemType indicates an invalid item type was found
	ErrBadItemType
	// ErrBadGearType indicates an invalid gear type was found
	ErrBadGearType
	// ErrBadClassType indicates an invalid class type was encountered
	ErrBadClassType
	// ErrBadSkillType indicates an invalid skill type was encountered
	ErrBadSkillType
	// ErrBadAttackSpeed indicates an invalid attack speed was encountered
	ErrBadAttackSpeed
	// ErrBadElement indicates an invalid element ID was encountered
	ErrBadElement
	// ErrBadPowderTier indicates an invalid powder tier was encountered
	ErrBadPowderTier
	// ErrBadConsumableType indicates an invalid consumable type was encountered
	ErrBadConsumableType
	// ErrBadEffectType indicates an invalid effect type was encountered
	ErrBadEffectType
	// ErrUnexpectedEndOfBytes indicates the decoder ran out of bytes to decode
	ErrUnexpectedEndOfBytes
	// ErrBadCodepoint indicates the decoder hit an invalid codepoint
	ErrBadCodepoint
)

// DecodeError represents errors that can occur during decoding
type DecodeError struct {
	Type    DecodeErrorType
	Details interface{}
}

// Error returns the error message for a DecodeError
func (e *DecodeError) Error() string {
	switch e.Type {
	case ErrNoStartBlockFound:
		return "No start block found"
	case ErrUnknownVersion:
		return fmt.Sprintf("Unknown version: %v", e.Details)
	case ErrStartReparse:
		return "Second start block found in data"
	case ErrUnknownBlock:
		return fmt.Sprintf("Unknown block ID: %v", e.Details)
	case ErrBadString:
		return "Decoder decoded a bad string"
	case ErrBadItemType:
		return fmt.Sprintf("Invalid item type id: %v", e.Details)
	case ErrBadGearType:
		return fmt.Sprintf("Invalid gear type id: %v", e.Details)
	case ErrBadClassType:
		return fmt.Sprintf("Invalid class type id: %v", e.Details)
	case ErrBadSkillType:
		return fmt.Sprintf("Invalid skill type id: %v", e.Details)
	case ErrBadAttackSpeed:
		return fmt.Sprintf("Invalid attack speed id: %v", e.Details)
	case ErrBadElement:
		return fmt.Sprintf("Invalid element id: %v", e.Details)
	case ErrBadPowderTier:
		return fmt.Sprintf("Invalid powder tier: %v", e.Details)
	case ErrBadConsumableType:
		return fmt.Sprintf("Invalid consumable type id: %v", e.Details)
	case ErrBadEffectType:
		return fmt.Sprintf("Invalid effect type: %v", e.Details)
	case ErrUnexpectedEndOfBytes:
		return "Unexpectedly hit end of bytestream while decoding"
	case ErrBadCodepoint:
		return fmt.Sprintf("Bad codepoint: %v", e.Details)
	default:
		return "Unknown decoding error"
	}
}

// DecoderError provides context about which block caused a decoding error
type DecoderError struct {
	ErrorData *DecodeError
	During    *DataBlockID // Optional: may be nil if block is unknown
}

// Error returns the error message for a DecoderError
func (e *DecoderError) Error() string {
	if e.During == nil {
		return fmt.Sprintf("Error while decoding unknown block. Type: %v, %v", e.ErrorData.Type, e.ErrorData.Details)
	}
	return fmt.Sprintf("Error while decoding block %d. Type: %v, %v", *e.During, e.ErrorData.Type, e.ErrorData.Details)
}
