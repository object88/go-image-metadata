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
type TagInitializer func(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Tag

type TagReader interface {
	GetReader() reader.Reader
	ReadIfd(ifdAddress uint32)
}
