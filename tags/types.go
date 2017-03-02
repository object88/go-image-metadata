package tags

import (
	"fmt"

	"github.com/object88/go-image-metadata/common"
	"github.com/object88/go-image-metadata/reader"
)

// Metadata is a single name-value pair.
type Metadata interface {
	fmt.Stringer
}

// TagID is a metadata identifier
type TagID uint16

// TagInitializer takes raw data and returns a Metadata
type TagInitializer func(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata

type TagReader interface {
	GetReader() reader.Reader
	ReadIfd(ifdAddress uint32)
}
