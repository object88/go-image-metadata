package reader

import (
	"encoding/binary"
	"io"
)

// BigEndianReader is a wrapper around `bytes.Reader` with respect towards
// handling big endian byte streams
type BigEndianReader struct {
	base
}

func CreateBigEndianReader(r io.ReadSeeker, baseOffset int64) Reader {
	return &BigEndianReader{
		base: base{r, baseOffset},
	}
}

func (r *BigEndianReader) ReadUint8FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xff
	if count > 0 {
		result[1] = data & 0xff00 >> 8
		if count > 1 {
			result[2] = data & 0xff0000 >> 16
			if count > 2 {
				result[3] = data & 0xff000000 >> 24
			}
		}
	}

	return result, nil
}

func (r *BigEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.base.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(t), nil
}

func (r *BigEndianReader) ReadUint16FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xffff
	if count > 0 {
		result[1] = data & 0xffff0000 >> 16
	}
	return result, nil
}

func (r *BigEndianReader) ReadUint32() (uint32, error) {
	t, err := readBytes(r.base.r, 4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(t), nil
}

func (r *BigEndianReader) ReadUint64() (uint64, error) {
	t, err := readBytes(r.base.r, 8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(t), nil
}
