package tiff

import (
	"fmt"
	"io"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/reader"
	"github.com/object88/go-image-metadata/tags"
)

func init() {
	metadata.RegisterHeaderCheck(CheckMotorolaHeader)
}

// MotorolaReader wraps a little-endian byte reader to expose Exif data
type MotorolaReader struct {
	r reader.Reader
}

// CheckMotorolaHeader peeks at the byte stream for the magic numbers to
// identify this as a TIFF image with big-endian encoding.
func CheckMotorolaHeader(r io.ReadSeeker) (metadata.ImageReader, error) {
	fmt.Printf("Checking motorola tiff header... ")
	cur, _ := r.Seek(0, io.SeekCurrent)
	b := []byte{0x00, 0x00, 0x00, 0x00}
	n, err := r.Read(b)
	if n != 4 || err != nil {
		return nil, err
	}

	// Read the magic number and endian check
	if b[0] != 0x4d || b[1] != 0x4d || b[2] != 0x00 || b[3] != 0x2a {
		fmt.Printf("got %#v; was wrong\n", b)
		return nil, nil
	}

	fmt.Printf("matched!\n")
	return &MotorolaReader{r: reader.CreateBigEndianReader(r, cur)}, nil
}

func (r *MotorolaReader) Read() map[uint16]tags.Tag {
	m := map[uint16]tags.Tag{}
	r.ReadPartial(&m)
	return m
}

func (r *MotorolaReader) ReadPartial(foundTags *map[uint16]tags.Tag) int64 {
	return 0
}
