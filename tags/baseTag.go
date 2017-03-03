package tags

// BaseTag is the common struct that other type-specific tags are composed with
type BaseTag struct {
	tagID TagID
}

// GetID returns the tag ID
func (b *BaseTag) GetID() TagID {
	return b.tagID
}

// GetName returns the coloquial name of the tag
func (b *BaseTag) GetName() string {
	tagBuilder := TagMap[uint16(b.tagID)]
	return tagBuilder.GetName()
}
