package reader

import (
	"errors"
	"fmt"
	"io"
)

type base struct {
	r io.ReadSeeker
}

func (r *base) Discard(count int) error {
	c := int64(count)
	_, err := r.r.Seek(c, io.SeekCurrent)
	if err != nil {
		panic(err)
	}
	fmt.Printf("; moved by %d bytes", count)
	return nil
}

func (r *base) GetReader() io.ReadSeeker {
	return r.r
}

func (r *base) ReadNullTerminatedString() (string, error) {
	start, err := r.r.Seek(0, io.SeekCurrent)
	if err != nil {
		return "", err
	}
	b := []byte{0x00}
	length := 0
	for {
		n, err := r.r.Read(b)
		if err != nil {
			return "", err
		}
		if n != 1 {
			return "", errors.New("Not enough byte")
		}
		if b[0] == '\x00' {
			// This is the end.
			break
		}
		length++
	}

	r.r.Seek(start, io.SeekStart)
	buf := make([]byte, length)

	n, err := r.r.Read(buf)
	if err != nil {
		return "", err
	}
	if n != length {
		return "", errors.New("Failed to re-read text\n")
	}

	return string(buf), nil
}

func (r *base) ReadTo() (bool, error) {
	fmt.Printf("; reading past image segment")
	b := []byte{0x00}
	passed := 0

	for {
		n, err := r.r.Read(b)
		if n != 1 || err != nil {
			return false, errors.New("Failed to read 1 byte")
		}
		passed++
		if b[0] == '\xff' {
			// Check the next byte; if it is non-0x00, we are done.
			n, err := r.r.Read(b)
			if n != 1 || err != nil {
				return false, errors.New("Failed to read 1 next byte")
			}
			if b[0] != '\x00' {
				// Found 0xffxx, where xx != 00
				r.r.Seek(-2, io.SeekCurrent)
				fmt.Printf("; after %d bytes, found non-escaped 0xff", passed)
				return true, nil
			}
		}
	}
}

func (r *base) ReadUint8() (uint8, error) {
	t, err := readBytes(r.r, 1)
	if err != nil {
		return 0, err
	}
	return t[0], nil
}

func readBytes(r io.Reader, size int) ([]byte, error) {
	t := make([]byte, size)
	bytesRead, err := r.Read(t)
	if err != nil {
		return nil, err
	} else if bytesRead != size {
		return nil, fmt.Errorf("Was only able to read %d of %d bytes", bytesRead, size)
	}
	return t, nil
}
