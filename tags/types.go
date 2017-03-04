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
type TagInitializer func(reader TagReader, foundTags *map[uint16]Tag, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error)

type TagReader interface {
	GetReader() reader.Reader
	ReadIfd(ifdAddress uint32, tags map[uint16]TagBuilder, foundTags *map[uint16]Tag)
}
