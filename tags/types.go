package tags

import "github.com/object88/go-image-metadata/common"

// TagInitializer takes raw data and returns a Metadata
type TagInitializer func(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata
