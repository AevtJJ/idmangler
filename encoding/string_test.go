package encoding

import (
	"bytes"
	"testing"
)

func TestStringRoundTrip(t *testing.T) {
	testCases := [][]byte{
		{},
		{0},
		{1, 2, 3},
		{255, 255},
		{255, 254},
		{255, 255, 0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 255, 254, 255, 255},
	}

	for _, tc := range testCases {
		encoded := EncodeString(tc)
		decoded, err := DecodeString(encoded)

		if err != nil {
			t.Errorf("Error decoding %v: %v", tc, err)
			continue
		}

		if !bytes.Equal(decoded, tc) {
			t.Errorf("Bytes mismatch after round trip. Expected %v, got %v", tc, decoded)
		}
	}
}

func TestEncodeDecodeChar(t *testing.T) {
	testCases := []struct {
		a byte
		b *byte
	}{
		{0, nil},
		{42, nil},
		{255, nil},
		{0, bytePtr(0)},
		{42, bytePtr(100)},
		{255, bytePtr(254)},
		{255, bytePtr(255)},
	}

	for _, tc := range testCases {
		char := EncodeChar(tc.a, tc.b)

		bytes, err := DecodeChar(char)
		if err != nil {
			t.Errorf("Error decoding char for %v, %v: %v", tc.a, tc.b, err)
			continue
		}

		if tc.b == nil {
			if len(bytes) != 1 || bytes[0] != tc.a {
				t.Errorf("Byte mismatch after single byte round trip. Expected [%d], got %v", tc.a, bytes)
			}
		} else {
			if len(bytes) != 2 || bytes[0] != tc.a || bytes[1] != *tc.b {
				t.Errorf("Byte mismatch after double byte round trip. Expected [%d, %d], got %v", tc.a, *tc.b, bytes)
			}
		}
	}
}

// Helper function to get a pointer to a byte
func bytePtr(b byte) *byte {
	return &b
}
