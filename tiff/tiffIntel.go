package tiff

import (
	"bufio"
	"fmt"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
)

func init() {
	metadata.RegisterHeaderCheck(CheckIntelHeader)
}

// IntelReader wraps a big-endian byte reader to expose Exif data
type IntelReader struct {
	r reader.Reader
}

func CheckIntelHeader(r *bufio.Reader) (common.ImageReader, error) {
	// b := make([]byte, 4)
	// n, err := r.Read(b)
	b, err := r.Peek(4)
	if err != nil {
		return nil, err
	}
	// if n != 4 {
	// 	return nil, errors.New("Did not read 4 bytes")
	// }

	// Read the magic number and endian check
	if b[0] != 0x49 || b[1] != 0x49 || b[2] != 0x2a || b[3] != 0x00 {
		fmt.Printf("Got %#v; was wrong\n", b)
		return nil, nil
	}

	fmt.Printf("Matched intel Tiff reader\n")
	return &IntelReader{r: reader.CreateLittleEndianReader(r)}, nil
}

func (r *IntelReader) Read() {
	// We know what's in the first 4 bytes, so we can skip past those.  The next
	// 4 bytes are the address of the first IFD
	r.r.Discard(4)
	_, err := r.r.ReadUint32()
	if err != nil {
		panic(fmt.Sprintf("FAILED to read address of 1st IFD: %s", err))
	}

	// Problem; we have a forward-reading buffer with a forward-and-backward
	// reading encoding.

	// for {
	// 	// Loop over all IFD
	// 	ifdAddress, err := r.r.ReadUint64()
	//
	// }
}
