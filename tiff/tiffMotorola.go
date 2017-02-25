package tiff

import (
	"fmt"
	"io"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
)

func init() {
	metadata.RegisterHeaderCheck(CheckMotorolaHeader)
}

// MotorolaReader wraps a little-endian byte reader to expose Exif data
type MotorolaReader struct {
	r reader.Reader
}

func CheckMotorolaHeader(r io.ReadSeeker) (common.ImageReader, error) {
	b := []byte{0x00, 0x00, 0x00, 0x00}
	n, err := r.Read(b)
	if n != 4 || err != nil {
		return nil, err
	}

	// Read the magic number and endian check
	if b[0] != 0x4d || b[1] != 0x4d || b[2] != 0x00 || b[3] != 0x2a {
		fmt.Printf("Got %s; was wrong\n", b)
		return nil, nil
	}

	fmt.Printf("Matched motorola Tiff reader\n")
	return &MotorolaReader{r: reader.CreateBigEndianReader(r)}, nil
}

func (r *MotorolaReader) Read() {}
