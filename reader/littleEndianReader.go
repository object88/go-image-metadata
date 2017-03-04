package reader

import (
	"encoding/binary"
	"io"
)

// LittleEndianReader is a wrapper around `bytes.Reader` with respect towards
// handling little endian byte streams
type LittleEndianReader struct {
	base
}

// CreateLittleEndianReader wraps an io.Reader with logic to read little-endian
// byte content.  Operations in the reader are relative to the provided
// baseOffset.
func CreateLittleEndianReader(r io.ReadSeeker, baseOffset int64) *LittleEndianReader {
	return &LittleEndianReader{
		base: base{r, baseOffset},
	}
}

// ReadUint8FromUint32 reads uint8s off the provided uint32, up to a maximum of
// count
func (r *LittleEndianReader) ReadUint8FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xff000000 >> 24
	if count > 1 {
		result[1] = data & 0xff0000 >> 16
		if count > 2 {
			result[2] = data & 0xff00 >> 8
			if count > 3 {
				result[3] = data & 0xff
			}
		}
	}

	return result, nil
}

// ReadUint16 reads 16 bits of unsigned data
func (r *LittleEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(t), nil
}

// ReadUint16FromUint32 reads uint16s off the provided uint32, up to a maximum
// of count
func (r *LittleEndianReader) ReadUint16FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xffff0000 >> 16
	if count > 1 {
		result[1] = data & 0xffff
	}
	return result, nil
}

// ReadUint32 reads 32 bits of unsigned data
func (r *LittleEndianReader) ReadUint32() (uint32, error) {
	t, err := readBytes(r.r, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(t), nil
}

// ReadUint64 reads 64 bits of unsigned data
func (r *LittleEndianReader) ReadUint64() (uint64, error) {
	t, err := readBytes(r.r, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(t), nil
}
