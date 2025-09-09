package encoding

import (
	"fmt"
)

// Error types
const (
	// ErrUnexpectedEndOfBytes indicates that the byte array ended unexpectedly
	ErrUnexpectedEndOfBytes = iota
	// ErrStartReparse indicates that a start block was encountered during reparsing
	ErrStartReparse
	// ErrUnknownBlock indicates that an unknown block ID was encountered
	ErrUnknownBlock
	// ErrNoBasevalueGiven indicates that no base value was provided for a stat
	ErrNoBasevalueGiven
	// ErrTooManyIdentifications indicates that there are too many identifications
	ErrTooManyIdentifications
	// ErrNoTypeGiven indicates that no item type was given
	ErrNoTypeGiven
	// ErrNoNameGiven indicates that no name was given
	ErrNoNameGiven
	// ErrInvalidVarInt indicates a problem with VarInt encoding/decoding
	ErrInvalidVarInt
	// ErrNoStartBlockFound indicates that no start block was found when expected
	ErrNoStartBlockFound
	// ErrUnknownVersion indicates an unknown encoding version
	ErrUnknownVersion
	// ErrBadItemType indicates an invalid item type
	ErrBadItemType
	// ErrNonAsciiString indicates that a string contains non-ASCII characters
	ErrNonAsciiString
	// ErrTooManyPowders indicates that there are too many powders
	ErrTooManyPowders
	// ErrBadElement indicates an invalid element type
	ErrBadElement
	// ErrBadPowderTier indicates an invalid powder tier
	ErrBadPowderTier
)

// DataBlockID is used for error reporting
type DataBlockID byte

// Encode errors
// EncodeErrorType represents the type of an encoding error
type EncodeErrorType int

// EncodeError represents an error that occurred during encoding
type EncodeError struct {
	Type    EncodeErrorType
	Details interface{}
}

// Error returns the error message for an encoding error
func (e *EncodeError) Error() string {
	switch e.Type {
	case ErrNoTypeGiven:
		return "No type given while encoding TypeData"
	case ErrNoNameGiven:
		return "No name given while encoding NameData"
	case ErrNoBasevalueGiven:
		return fmt.Sprintf("No base value given for stat with kind %v", e.Details)
	case ErrTooManyIdentifications:
		return "Too many identifications (maximum is 255)"
	case ErrNonAsciiString:
		return "String contains non-ASCII characters"
	case ErrTooManyPowders:
		return "Too many powders (maximum is 6)"
	case ErrBadElement:
		return fmt.Sprintf("Invalid element: %v", e.Details)
	case ErrBadPowderTier:
		return fmt.Sprintf("Invalid powder tier: %v", e.Details)
	case ErrBadItemType:
		return fmt.Sprintf("Invalid item type: %v", e.Details)
	default:
		return fmt.Sprintf("Unknown encoding error: %d", e.Type)
	}
}

// Decode errors
// DecodeErrorType represents the type of a decoding error
type DecodeErrorType int

// DecodeError represents an error that occurred during decoding
type DecodeError struct {
	Type    DecodeErrorType
	Details interface{}
}

// Error returns the error message for a decoding error
func (e *DecodeError) Error() string {
	switch e.Type {
	case ErrUnexpectedEndOfBytes:
		return "Unexpected end of bytes"
	case ErrStartReparse:
		return "Start block encountered during reparsing"
	case ErrUnknownBlock:
		return fmt.Sprintf("Unknown block: %v", e.Details)
	case ErrInvalidVarInt:
		return fmt.Sprintf("Invalid VarInt: %v", e.Details)
	case ErrNoStartBlockFound:
		return "No start block found"
	case ErrUnknownVersion:
		return fmt.Sprintf("Unknown version: %v", e.Details)
	case ErrBadItemType:
		return fmt.Sprintf("Invalid item type: %v", e.Details)
	case ErrBadElement:
		return fmt.Sprintf("Invalid element: %v", e.Details)
	case ErrBadPowderTier:
		return fmt.Sprintf("Invalid powder tier: %v", e.Details)
	default:
		return fmt.Sprintf("Unknown decoding error: %d", e.Type)
	}
}

// DecoderError wraps a decode error with context about which block was being decoded
type DecoderError struct {
	ErrorData *DecodeError
	During    *DataBlockID
}

// Error returns the error message for a decoder error
func (e *DecoderError) Error() string {
	if e.During != nil {
		return fmt.Sprintf("Error while decoding block %d: %s", *e.During, e.ErrorData.Error())
	}
	return e.ErrorData.Error()
}
