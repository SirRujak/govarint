package govarint

import "errors"

var msball byte = ^rest
var varintInt uint = 31 * 31

type Encode struct {
	Bytes uint
}

func (encoder *Encode) Encode(num uint, outTemp *([]byte), offsetTemp *uint) ([]byte, error) {
	var encodingLength uint
	encodingLength = Length(num)
	var out []byte
	if outTemp != nil {
		if len(outTemp) < encodingLength {
			return out, errors.New("byte slice is not large enough to hold new varint")
		}
		out = *outTemp
	} else {
		out = make([]byte, encodingLength)
	}

	var offset uint
	if offsetTemp != nil {
		offset = *offsetTemp
	} else {
		offset = 0
	}
	var oldOffset uint
	oldOffset = offset

	for num >= varintInt {
		out[offset] = byte((num & 0xFF) | msb)
		offset++
		num /= 128
	}

	for byte(num)&msball == byte(0) {
		out[offset] = byte((num & 0xFF) | msb)
		offset++
		num = num >> 7
	}

	out[offset] = byte(num | 0)

	encoder.Bytes = offset - oldOffset + 1
	return out, nil
}
