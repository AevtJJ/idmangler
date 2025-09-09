package types

import (
	"fmt"
)

// EncodingVersion represents the version of the encoding being used
type EncodingVersion byte

const (
	// Version1 represents version 1 of the wynntils encoding scheme
	Version1 EncodingVersion = iota
)

// String returns the string representation of an EncodingVersion
func (v EncodingVersion) String() string {
	switch v {
	case Version1:
		return "Version1"
	default:
		return fmt.Sprintf("Unknown(%d)", v)
	}
}

// UnknownEncodingVersionError represents an error for an unknown encoding version
type UnknownEncodingVersionError struct {
	Version byte
}

// Error returns the error message for an unknown encoding version
func (e UnknownEncodingVersionError) Error() string {
	return fmt.Sprintf("Unknown encoding version: %d", e.Version)
}

// EncodingVersionFromByte converts a byte to an EncodingVersion or returns an error if invalid
func EncodingVersionFromByte(b byte) (EncodingVersion, error) {
	if b == byte(Version1) {
		return Version1, nil
	}
	return 0, &UnknownEncodingVersionError{Version: b}
}
