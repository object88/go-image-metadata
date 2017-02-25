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

func CreateBigEndianReader(r io.ReadSeeker) Reader {
	return &BigEndianReader{
		base: base{r},
	}
}

func (r *BigEndianReader) ReadUint16() (uint16, error) {
	t, err := readBytes(r.base.r, 2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(t), nil
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
