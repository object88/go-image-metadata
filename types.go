package metadata

import (
	"io"

	"github.com/object88/go-image-metadata/common"
)

type marker uint16

type markerSegment struct {
	marker marker
	size   uint16
	data   []byte
}

// CheckHeader examines the first bytes of a byte stream for magic numbers, or
// other prefixes, in order to validate that the implementor can accept the byte
// stream.  If the byte stream conforms to the shape readable by the
// implementor, an ImageReader should be returned.  An error should only be
// returned if there is a problem reading the byte stream, not in the case of
// non-conformity.
type CheckHeader func(reader io.ReadSeeker) (common.ImageReader, error)
