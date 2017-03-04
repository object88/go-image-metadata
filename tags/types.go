package tags

import (
	"fmt"

	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
)

// Tag is a single name-value pair.
type Tag interface {
	fmt.Stringer
}

// TagID is a tag identifier
type TagID uint16

// TagInitializer takes raw data and returns a Tag
type TagInitializer func(reader TagReader, foundTags *map[uint16]Tag, name string, raw *RawTagData) (Tag, bool, error)

// TagReader implementations will read over an IFD in a image file, creating
// Tags and putting them in foundTags.
type TagReader interface {
	GetReader() reader.Reader
	ReadIfd(ifdAddress uint32, tags []*map[uint16]TagBuilder, foundTags *map[uint16]Tag)
}

// RawTagData contains the data as read from the images byte stream
type RawTagData struct {
	// Tag is the id
	Tag TagID

	// Format describes the data type
	Format common.DataFormat

	// Count indicates the number of entries to be read
	Count uint32

	// Data either contains the whole of the data (if 4 bytes of less) or a
	// pointer to the actual data location
	Data uint32
}
