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

// CreateBigEndianReader wraps an io.Reader with logic to read big-endian byte
// content.  Operations in the reader are relative to the provided baseOffset.
func CreateBigEndianReader(r io.ReadSeeker, baseOffset int64) Reader {
	return &BigEndianReader{
		base: base{r, baseOffset},
	}
}

// ReadUint8FromUint32 reads uint8s off the provided uint32, up to a maximum of
// count
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

// ReadUint16 reads 16 bits of unsigned data
func (r *BigEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.base.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(t), nil
}

// ReadUint16FromUint32 reads uint16s off the provided uint32, up to a maximum
// of count
func (r *BigEndianReader) ReadUint16FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xffff
	if count > 0 {
		result[1] = data & 0xffff0000 >> 16
	}
	return result, nil
}

// ReadUint32 reads 32 bits of unsigned data
func (r *BigEndianReader) ReadUint32() (uint32, error) {
	t, err := readBytes(r.base.r, 4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(t), nil
}

// ReadUint64 reads 64 bits of unsigned data
func (r *BigEndianReader) ReadUint64() (uint64, error) {
	t, err := readBytes(r.base.r, 8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(t), nil
}
