package jfif

import (
	"bufio"
	"fmt"

	"github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
)

func init() {
	metadata.RegisterHeaderCheck(CheckHeader)
}

type Reader struct {
	r reader.Reader
}

// CheckHeader checks the byte stream to see if it contains a JFIF
func CheckHeader(r *bufio.Reader) (common.ImageReader, error) {
	// b := make([]byte, 2)
	// n, err := r.Read(b)
	b, err := r.Peek(2)
	if err != nil {
		return nil, err
	}
	// if n != 2 {
	// 	return nil, errors.New("Did not read 2 bytes")
	// }

	if b[0] != 0xff || b[1] != 0xd8 {
		return nil, nil
	}
	fmt.Printf("Matched Jfif reader\n")
	return &Reader{r: reader.CreateBigEndianReader(r)}, nil
}

func (r *Reader) Read() {
	// Loop over marker segments
	for {
		m, e := r.r.ReadUint16()
		if e != nil {
			panic(fmt.Errorf("NOOOOO: %s", e))
		}

		fmt.Printf("\n0x%04x", m)

		m1 := marker(m)
		if m1 == eoi {
			// We have reached the end of the file.
			fmt.Printf("; DONE\n")
			break
		}

		if m&0xffe0 == 0xffe0 {
			// We have an appN segment.
			r.readAppnSegment()
		} else if m1 == soi || m^0xffd0>>3 == 0 {
			// Restart: 0xffd0-0xffd7; nothing to process.
			fmt.Printf("; got restart")
			continue
		} else if m1 == sos {
			// This is the beginning of the image data.  We want to scan past all
			// this, but we don't have a length.
			r.movePastImageSegment()
		} else {
			r.moveToNextSegment()
		}
	}
}

func (r *Reader) readAppnSegment() {
	fmt.Printf("; app segment")
	remaining, err := r.r.ReadUint16()
	if err != nil {
		panic(err)
	}

	remaining -= 2
	fmt.Printf("; will read %d bytes", remaining)

	id, err := r.r.ReadNullTerminatedString()
	if err != nil {
		panic(err)
	}
	fmt.Printf("; found identifier: '%s' (%d)", id, len(id))

	remaining -= uint16(len(id))

	// Need to check the type of app segment by the null-terminated string, then
	// act appropriately.
	switch id {
	case "Exif\x00":
		fmt.Printf(" **WOOOO**\n")
		// The `Exif` string is double-null terminated:
		// https://www.media.mit.edu/pia/Research/deepview/exif.html
		r.r.Discard(1)
		remaining--
		r1, err := metadata.ReadHeader(r.r.GetReader())
		if err != nil {
			panic("NOPE")
		}
		r1.Read()

		r.r.Discard(int(remaining))
	default:
		r.r.Discard(int(remaining))
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
		fmt.Printf("; FF %d bytes", s0)
		r.r.Discard(int(s0))
	}
}

func (r *Reader) movePastImageSegment() {
	// There is no length in the SOS marker segment; it is just a stream of bytes
	// until we encounter 0xFF which is not immediately followed by 0x00
	// http://stackoverflow.com/questions/26715684/parsing-jpeg-sos-marker
	r.r.ReadTo()
}
