package reader

import (
	"bufio"
	"encoding/binary"
)

type LittleEndianReader struct {
	base
}

func CreateLittleEndianReader(r *bufio.Reader) *LittleEndianReader {
	return &LittleEndianReader{
		base: base{r},
	}
}

func (r *LittleEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(t), nil
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
