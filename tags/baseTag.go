package tags

// BaseTag is the common struct that other type-specific tags are composed with
type BaseTag struct {
	name  string
	tagID TagID
}

// GetID returns the tag ID
func (b *BaseTag) GetID() TagID {
	return b.tagID
}

// GetName returns the coloquial name of the tag
func (b *BaseTag) GetName() string {
	return b.name
}
