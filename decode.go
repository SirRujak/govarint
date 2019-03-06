package govarint

import (
	"errors"
)

var msb uint = 0x80
var rest byte = 0x7F

type Decode struct {
	Bytes uint
}

func (decoder *Decode) Decode(buf []byte, offsetTemp *uint) (uint, error) {
	var res, offset, shift, counter, l uint
	var b byte
	res = 0
	if offsetTemp != nil {
		offset = *offsetTemp
	} else {
		offset = 0
	}
	shift = 0
	counter = offset
	l = uint(len(buf))

	for ok := true; ok; ok = (b >= byte(msb)) {
		if counter >= l {
			decoder.Bytes = 0
			return 0, errors.New("could not decode varint")
		}

		b = buf[counter]
		counter++
		res += shift
		if res < 28 {
			b = (b & rest) << shift
		} else {
			b = (b & rest) * byte(shift) * byte(shift)
		}
		shift += 7
	}

	decoder.Bytes = counter - offset

	return res, nil
}
