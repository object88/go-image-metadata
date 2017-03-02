package reader

import (
	"encoding/binary"
	"io"
)

type LittleEndianReader struct {
	base
}

func CreateLittleEndianReader(r io.ReadSeeker, baseOffset int64) *LittleEndianReader {
	return &LittleEndianReader{
		base: base{r, baseOffset},
	}
}

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

func (r *LittleEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(t), nil
}

func (r *LittleEndianReader) ReadUint16FromUint32(count, data uint32) ([]uint32, error) {
	result := make([]uint32, count)
	result[0] = data & 0xffff0000 >> 16
	if count > 1 {
		result[1] = data & 0xffff
	}
	return result, nil
}

func (r *LittleEndianReader) ReadUint32() (uint32, error) {
	t, err := readBytes(r.r, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(t), nil
}

func (r *LittleEndianReader) ReadUint64() (uint64, error) {
	t, err := readBytes(r.r, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(t), nil
}
