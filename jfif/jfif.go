package jfif

import (
	"fmt"
	"io"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/reader"
	"github.com/object88/go-image-metadata/tags"
)

func init() {
	metadata.RegisterHeaderCheck(CheckHeader)
}

// Reader understands a Jfif byte stream
type Reader struct {
	r reader.Reader
}

// CheckHeader checks the byte stream to see if it contains a JFIF
func CheckHeader(r io.ReadSeeker) (metadata.ImageReader, error) {
	fmt.Printf("Checking jfif header... ")
	cur, _ := r.Seek(0, io.SeekCurrent)
	b := []byte{0x00, 0x00}
	n, err := r.Read(b)
	if n != 2 || err != nil {
		return nil, err
	}

	if b[0] != 0xff || b[1] != 0xd8 {
		fmt.Printf("got %#v; was wrong\n", b)
		return nil, nil
	}
	fmt.Printf("matched\n")
	return &Reader{r: reader.CreateBigEndianReader(r, cur)}, nil
}

func (r *Reader) Read() map[uint16]tags.Tag {
	m := map[uint16]tags.Tag{}
	r.ReadPartial(&m)
	return m
}

func (r *Reader) ReadPartial(foundTags *map[uint16]tags.Tag) int64 {
	// Loop over marker segments
	for {
		m, e := r.r.ReadUint16()
		if e != nil {
			panic(fmt.Errorf("NOOOOO: %s", e))
		}

		fmt.Printf("0x%04x; ", m)

		m1 := marker(m)
		if m1 == eoi {
			// We have reached the end of the file.
			fmt.Printf("DONE\n")
			break
		}

		if m&0xffe0 == 0xffe0 {
			// We have an appN segment.
			r.readAppnSegment(foundTags)
		} else if m1 == soi || m^0xffd0>>3 == 0 {
			// Restart: 0xffd0-0xffd7; nothing to process.
			fmt.Printf("got restart\n")
			continue
		} else if m1 == sos {
			// This is the beginning of the image data.  We want to scan past all
			// this, but we don't have a length.
			r.movePastImageSegment()
		} else {
			r.moveToNextSegment()
		}
	}
	return 0
}

func (r *Reader) readAppnSegment(foundTags *map[uint16]tags.Tag) {
	fmt.Printf("app segment")
	rem, err := r.r.ReadUint16()
	if err != nil {
		panic(err)
	}

	remaining := int64(rem) - 2
	fmt.Printf("; will read %d bytes", remaining)

	id, err := r.r.ReadNullTerminatedString()
	if err != nil {
		panic(err)
	}
	fmt.Printf("; found identifier: '%s' (%d)", id, len(id))

	remaining -= int64(len(id))

	// Need to check the type of app segment by the null-terminated string, then
	// act appropriately.
	switch id {
	case "Exif":
		fmt.Printf(" **WOOOO**\n")
		// The `Exif` string is double-null terminated:
		// https://www.media.mit.edu/pia/Research/deepview/exif.html
		r.r.Discard(2)
		remaining -= 2
		r1, err := metadata.ReadHeader(r.r.GetReader())
		if err != nil {
			panic("NOPE")
		}
		consumed := r1.ReadPartial(foundTags)
		remaining -= consumed

		r.r.Discard(int64(remaining))
	default:
		r.r.Discard(int64(remaining))
	}
}

func (r *Reader) moveToNextSegment() {
	// Ignore this segment.  Need to read the variable length, and scan past it.
	s, err := r.r.ReadUint16()
	if err != nil {
		panic(err)
	}

	s0 := s - 2
	if s0 > 0 {
		r.r.Discard(int64(s0))
	}
}

func (r *Reader) movePastImageSegment() {
	// There is no length in the SOS marker segment; it is just a stream of bytes
	// until we encounter 0xFF which is not immediately followed by 0x00
	// http://stackoverflow.com/questions/26715684/parsing-jpeg-sos-marker
	r.r.ReadTo()
}
