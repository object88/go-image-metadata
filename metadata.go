package metadata

import (
	"errors"
	"fmt"
	"io"

	"github.com/object88/go-image-metadata/common"
)

var readers []CheckHeader

func RegisterHeaderCheck(fn CheckHeader) {
	readers = append(readers, fn)
}

func ReadHeader(reader io.ReadSeeker) (common.ImageReader, error) {
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
