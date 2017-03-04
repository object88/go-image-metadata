package metadata

import (
	"errors"
	"fmt"
	"io"
)

var readers []CheckHeader

// RegisterHeaderCheck adds a CheckHeader method to the internal collection,
// to handle more image formats
func RegisterHeaderCheck(fn CheckHeader) {
	readers = append(readers, fn)
}

// ReadHeader loops over the collection of HeaderCheck methods to determine
// which ImageReader implementation to use
func ReadHeader(reader io.ReadSeeker) (ImageReader, error) {
	start, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	for _, f := range readers {
		_, err = reader.Seek(start, io.SeekStart)
		if err != nil {
			return nil, err
		}
		r, err := f(reader)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return nil, err
		}
		if r != nil {
			fmt.Printf("OK\n")
			return r, nil
		}
	}

	return nil, errors.New("Unknown file format")
}
