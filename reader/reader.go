package reader

import "bufio"

// Reader is a byte reader whose implementations is endian-aware
type Reader interface {
	// Discard fast-forwards over `count` bytes, discarding their contents
	Discard(count int) error

	GetReader() *bufio.Reader

	// ReadNullTerminatedString reads a series of bytes, until it encounters
	// '\000', and returns a string.
	ReadNullTerminatedString() (string, error)

	ReadTo() (bool, error)

	// ReadUint8 reads an unsigned 8-bit value
	ReadUint8() (uint8, error)

	// ReadUint16 reads an unsigned 16-bit value
	ReadUint16() (uint16, error)

	// ReadUint32 reads an unsigned 32-bit value
	ReadUint32() (uint32, error)

	// ReadUint64 reads an unsigned 64-bit value
	ReadUint64() (uint64, error)
}
