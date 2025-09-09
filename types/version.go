package types

import "fmt"

// EncodingVersion represents the version of the encoding format
type EncodingVersion byte

const (
	// Version1 is the initial encoding version used in Wynntils
	Version1 EncodingVersion = 1
)

// String returns a string representation of the encoding version
func (v EncodingVersion) String() string {
	switch v {
	case Version1:
		return "Version1"
	default:
		return fmt.Sprintf("Unknown(%d)", v)
	}
}

// EncodingVersionFromByte converts a byte to an EncodingVersion or returns an error if invalid
func EncodingVersionFromByte(b byte) (EncodingVersion, error) {
	switch b {
	case 1, 0:
		return Version1, nil
	default:
		return 0, fmt.Errorf("unknown encoding version: %d", b)
	}
}
