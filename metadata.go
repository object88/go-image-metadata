package metadata

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/object88/go-image-metadata/common"
)

var readers []CheckHeader

func RegisterHeaderCheck(fn CheckHeader) {
	readers = append(readers, fn)
}

func ReadHeader(reader *bufio.Reader) (common.ImageReader, error) {
	// b := bufio.NewReader(reader)
	for _, f := range readers {
		fmt.Printf("Checking reader %#v... ", f)
		r, err := f(reader)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return nil, err
		}
		if r != nil {
			fmt.Printf("OK\n")
			return r, nil
		}
		fmt.Printf("\n")
	}

	return nil, errors.New("Unknown file format")
}
