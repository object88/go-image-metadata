package tiff

import (
	"fmt"
	"io"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
	"github.com/object88/go-image-metadata/tags"
)

func init() {
	metadata.RegisterHeaderCheck(CheckIntelHeader)
}

// IntelReader wraps a big-endian byte reader to expose Exif data
type IntelReader struct {
	r reader.Reader
}

// CheckIntelHeader peeks at the byte stream for the magic numbers to identify
// this as a TIFF image with little-endian encoding.
func CheckIntelHeader(r io.ReadSeeker) (metadata.ImageReader, error) {
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

func (r *IntelReader) Read() map[uint16]tags.Tag {
	m := map[uint16]tags.Tag{}
	r.ReadPartial(&m)
	return m
}

func (r *IntelReader) ReadPartial(foundTags *map[uint16]tags.Tag) int64 {
	// We have already read the first 4 bytes from the header.
	start, _ := r.r.GetReader().Seek(0, io.SeekCurrent)
	start -= 4

	// The next 4 bytes is the address of the first IFD
	ifdAddress, err := r.r.ReadUint32()
	if err != nil {
		panic(fmt.Sprintf("FAILED to read address of 1st IFD: %s", err))
	}

	r.ReadIfd(ifdAddress, []*map[uint16]tags.TagBuilder{&tags.TagMap}, foundTags)

	cur, _ := r.r.GetReader().Seek(0, io.SeekCurrent)
	return cur - start
}

func (r *IntelReader) GetReader() reader.Reader {
	return r.r
}

func (r *IntelReader) ReadIfd(ifdAddress uint32, tagMaps []*map[uint16]tags.TagBuilder, foundTags *map[uint16]tags.Tag) {
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
			format := common.DataFormat(f)
			tagID := tags.TagID(t)
			raw := &tags.RawTagData{Tag: tagID, Format: format, Count: c, Data: d}

			matched := false
			for _, tagMap := range tagMaps {
				tag, ok := (*tagMap)[t]
				if !ok {
					continue
				}

				// fmt.Printf("%d-%d: 0x%04x, %s, 0x%08x, 0x%08x\n", ifdN, i, t, format, c, d)
				initializer := tag.GetInitializer()
				m, ok, err := initializer(r, foundTags, tag.GetName(), raw)
				if err != nil {
					continue
				}
				if !ok {
					continue
				}

				if m != nil {
					fmt.Printf("%d-%d: %s\n", ifdN, i, m)
					(*foundTags)[t] = m
				}

				matched = true
				break
			}

			if !matched {
				// Unknown tag!
				fmt.Printf("%d-%d: unknown: 0x%04x, %s, 0x%08x, 0x%08x\n", ifdN, i, t, format, c, d)
			}
		}

		var ifdReadErr error
		ifdAddress, ifdReadErr = r.r.ReadUint32()
		if ifdReadErr != nil {
			return
		}
		if ifdAddress == 0 {
			fmt.Printf("End of IFD\n")
			break
		}
	}
}
