package metadata

import "github.com/object88/go-image-metadata/tags"

// ImageReader reads the image tags and stuff.
type ImageReader interface {
	Read() map[uint16]tags.Tag
	ReadPartial(foundTags *map[uint16]tags.Tag) int64
}
