package encoding

import (
	"testing"
)

func TestVarIntRoundTrip(t *testing.T) {
	testCases := []int64{
		0,
		1,
		-1,
		42,
		-42,
		1000,
		-1000,
		32767,                // i16::MAX
		-32768,               // i16::MIN
		2147483647,           // i32::MAX
		-2147483648,          // i32::MIN
		9223372036854775807,  // i64::MAX
		-9223372036854775808, // i64::MIN
	}

	for _, tc := range testCases {
		bytes := EncodeVarInt(tc)
		decoded, bytesUsed, err := DecodeVarInt(bytes)

		if err != nil {
			t.Errorf("Error decoding %d: %v", tc, err)
		}

		if bytesUsed != len(bytes) {
			t.Errorf("Incorrect number of bytes used. Expected %d, got %d", len(bytes), bytesUsed)
		}

		if decoded != tc {
			t.Errorf("Value mismatch after round trip. Expected %d, got %d", tc, decoded)
		}
	}
}
