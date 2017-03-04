package tags

import "github.com/object88/go-image-metadata/common"

// BaseTag is the common struct that other type-specific tags are composed with
type BaseTag struct {
	name    string
	tagID   TagID
	tagType common.DataFormat
}

// GetID returns the tag ID
func (b *BaseTag) GetID() TagID {
	return b.tagID
}

// GetName returns the coloquial name of the tag
func (b *BaseTag) GetName() string {
	return b.name
}

// GetType returns the type of the tag
func (b *BaseTag) GetType() common.DataFormat {
	return b.tagType
}
