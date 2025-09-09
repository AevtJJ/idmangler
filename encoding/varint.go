package encoding

// EncodeVarInt encodes an integer of variable size into bytes using zigzag encoding
// This format is used by wynntils for efficient integer representation
func EncodeVarInt(value int64) []byte {
	// zigzag encoding to remove sign bit
	zigzagged := uint64((value << 1) ^ (value >> 63))

	// Calculate number of bytes needed (7 bits per byte)
	// Highest bit is used to indicate end of encoding
	numBytes := 1
	temp := zigzagged >> 7
	for temp != 0 {
		numBytes++
		temp >>= 7
	}

	outBytes := make([]byte, numBytes)
	for i := 0; i < numBytes; i++ {
		next := byte(zigzagged>>uint(7*i)) & 0x7F

		// Set high bit for all bytes except the last one
		if i < numBytes-1 {
			next |= 0x80 // 10000000
		}

		outBytes[i] = next
	}

	return outBytes
}

// DecodeVarInt decodes a variable sized integer from a byte slice
// Returns the decoded value, number of bytes consumed, and any error
func DecodeVarInt(bytes []byte) (int64, int, error) {
	if len(bytes) == 0 {
		return 0, 0, &DecodeError{Type: ErrUnexpectedEndOfBytes}
	}

	var value int64
	var i int

	for i = 0; i < len(bytes); i++ {
		b := bytes[i]

		// Add the bits from this byte (sans continuation bit) to the value
		value |= int64(b&0x7F) << uint(7*i)

		// If high bit is not set, this is the last byte
		if (b & 0x80) == 0 {
			i++ // include this byte in the count
			break
		}
	}

	// Convert from zigzag encoding back to signed
	value = int64((uint64(value) >> 1)) ^ -(value & 1)

	return value, i, nil
}

// DecodeVarIntFromIterator decodes a variable-length encoded integer from a byte iterator
func DecodeVarIntFromIterator(nextByte func() (byte, error)) (int64, error) {
	var value int64
	var data []byte

	for {
		b, err := nextByte()
		if err != nil {
			return 0, &DecodeError{Type: ErrUnexpectedEndOfBytes}
		}

		data = append(data, b)

		if (b & 0b10000000) == 0 {
			break
		}
	}

	for i, b := range data {
		value |= int64(b&0b01111111) << (7 * i)
	}

	// Convert from zigzag encoding back to signed integer
	return (value >> 1) ^ -(value & 1), nil
}
