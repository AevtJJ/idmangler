package encoding

import "fmt"

// Constants defining Unicode private use areas used for encoding
const (
	// AreaA is the start of the supplementary private use area A
	AreaA uint32 = 0x0F0000
	// AreaB is the start of the supplementary private use area B
	AreaB uint32 = 0x100000
)

// BadCodepointError represents an error for an invalid codepoint during decoding
type BadCodepointError struct {
	Codepoint uint32
}

// Error returns the error message for a bad codepoint
func (e *BadCodepointError) Error() string {
	return fmt.Sprintf("Invalid codepoint: %06X", e.Codepoint)
}

// EncodeString encodes bytes into a string using the Wynntils byte encoding scheme
// https://github.com/Wynntils/Wynntils/blob/main/common/src/main/java/com/wynntils/utils/EncodedByteBuffer.java#L87
func EncodeString(data []byte) string {
	out := make([]rune, 0, (len(data)+len(data)%2)*2)

	for i := 0; i < len(data); i += 2 {
		if i+1 < len(data) {
			// Two bytes available
			out = append(out, EncodeChar(data[i], &data[i+1]))
		} else {
			// Only one byte available
			out = append(out, EncodeChar(data[i], nil))
		}
	}

	return string(out)
}

// EncodeChar encodes one or two bytes into a single character using the private use area encoding
func EncodeChar(a byte, b *byte) rune {
	if a == 0xFF && b != nil && *b >= 254 {
		return rune(AreaB + uint32(*b-254))
	} else if b != nil {
		return rune(AreaA + (uint32(a) << 8) + uint32(*b))
	} else {
		return rune(AreaB + (uint32(a) << 8) + 0xEE)
	}
}

// DecodeString decodes a Wynntils private area encoded string into bytes
// https://github.com/Wynntils/Wynntils/blob/main/common/src/main/java/com/wynntils/utils/EncodedByteBuffer.java#L33
func DecodeString(data string) ([]byte, error) {
	out := make([]byte, 0, len(data)*2)

	for _, c := range data {
		bytes, err := DecodeChar(c)
		if err != nil {
			return nil, err
		}
		out = append(out, bytes...)
	}

	return out, nil
}

// DecodeChar decodes a single character from the Wynntils private use area encoding scheme
// Returns 1 or 2 bytes or an error if the codepoint is not valid within the encoding scheme
func DecodeChar(c rune) ([]byte, error) {
	n := uint32(c)

	if n < AreaA || n > (AreaB+0xFFFF) {
		return nil, &BadCodepointError{Codepoint: n}
	}

	// Special cases
	if n >= AreaB {
		// Single byte
		if n&0xFF == 0xEE {
			return []byte{byte((n & 0xFF00) >> 8)}, nil
		}

		// Two bytes
		return []byte{255, byte(254 + (n & 0xFF))}, nil
	}

	// Normal case - two bytes
	return []byte{byte((n & 0xFF00) >> 8), byte(n & 0x00FF)}, nil
}
