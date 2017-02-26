package tiff

import (
	"fmt"
	"io"

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

func CheckIntelHeader(r io.ReadSeeker) (common.ImageReader, error) {
	fmt.Printf("Checking intel tiff header... ")

	cur, _ := r.Seek(0, io.SeekCurrent)

	b := []byte{0x00, 0x00, 0x00, 0x00}
	n, err := r.Read(b)
	if n != 4 || err != nil {
		return nil, err
	}

	// Read the magic number and endian check
	if b[0] != 0x49 || b[1] != 0x49 || b[2] != 0x2a || b[3] != 0x00 {
		fmt.Printf("got %#v; was wrong\n", b)
		return nil, nil
	}

	fmt.Printf("matched!\n")
	return &IntelReader{r: reader.CreateLittleEndianReader(r, cur)}, nil
}

func (r *IntelReader) Read() int64 {
	// We have already read the first 4 bytes from the header.
	start, _ := r.r.GetReader().Seek(0, io.SeekCurrent)
	start -= 4

	// The next 4 bytes is the address of the first IFD
	ifdAddress, err := r.r.ReadUint32()
	if err != nil {
		panic(fmt.Sprintf("FAILED to read address of 1st IFD: %s", err))
	}

	ifdN := -1
	for {
		// Loop over all IFD
		ifdN++
		fmt.Printf("Moving to IFD #%d at 0x%04x\n", ifdN, ifdAddress)
		r.r.SeekTo(int64(ifdAddress))

		count, _ := r.r.ReadUint16()
		for i := uint16(0); i < count; i++ {
			t, _ := r.r.ReadUint16()
			f, _ := r.r.ReadUint16()
			c, _ := r.r.ReadUint32()
			d, _ := r.r.ReadUint32()

			format := dataFormat(f)
			if c*d > 4 {
				// d is a pointer.
				if format == asciiString {
					cur := r.r.GetCurrentOffset()
					r.r.SeekTo(int64(d))
					s, _ := r.r.ReadNullTerminatedString()
					fmt.Printf("%d-%d: 0x%04x, %s, 0x%08x, 0x%08x: %s\n", ifdN, i, t, format, c, d, s)
					r.r.SeekTo(cur)
				} else {
					fmt.Printf("%d-%d: 0x%04x, %s, 0x%08x, 0x%08x\n", ifdN, i, t, format, c, d)
				}
			} else {
				fmt.Printf("%d-%d: 0x%04x, %s, 0x%08x, 0x%08x\n", ifdN, i, t, format, c, d)
			}
		}

		ifdAddress, err = r.r.ReadUint32()
		if err != nil {
			return 0
		}
		if ifdAddress == 0 {
			fmt.Printf("End of IFD\n")
			break
		}
	}

	cur, _ := r.r.GetReader().Seek(0, io.SeekCurrent)
	return cur - start
}
