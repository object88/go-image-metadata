package reader

import (
	"bufio"
	"fmt"
	"io"
)

type base struct {
	r *bufio.Reader
}

func (r *base) Discard(count int) error {
	passed, err := r.r.Discard(count)
	if err != nil {
		panic(err)
	}
	if count != passed {
		panic(fmt.Errorf("Expected to pass %d bytes; only moved %d", count, passed))
	}
	fmt.Printf("; moved by %d bytes", count)
	return nil
}

func (r *base) GetReader() *bufio.Reader {
	return r.r
}

func (r *base) ReadNullTerminatedString() (string, error) {
	b, err := r.r.ReadBytes('\x00')
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *base) ReadTo() (bool, error) {
	passed := 0
	for {
		l, err := r.r.ReadSlice('\xff')
		if err == bufio.ErrBufferFull {
			// Filled up the buffer without finding the character.
			fmt.Printf("Buffer filled...\n")
			continue
		} else if err == io.EOF {
			// Ran out of bytes entirely.
			fmt.Printf("EOF\n")
			return false, nil
		} else if err != nil {
			fmt.Printf("Other error: %s", err)
			return false, err
		}
		passed += len(l)
		next, err2 := r.r.Peek(1)
		if err2 != nil {
			if err2 == bufio.ErrBufferFull {
				fmt.Printf("Got 0xff and buffer full when peaking")
				return false, nil
			} else if err2 == io.EOF {
				fmt.Printf("Got 0xff and EOF")
				return false, nil
			}
			fmt.Printf("Got 0xff and other error: %s\n", err2)
			return false, err2
		}
		if next[0] != '\x00' {
			// Found 0xffxx, where xx != 00
			r.r.UnreadByte()
			fmt.Printf("; after %d bytes, found non-escaped 0xff", passed)
			return true, nil
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
