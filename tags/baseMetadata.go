package tags

type BaseMetadata struct {
	tagID TagID
}

func (b *BaseMetadata) GetID() TagID {
	return b.tagID
}

func (b *BaseMetadata) GetName() string {
	tagBuilder := TagMap[uint16(b.tagID)]
	return tagBuilder.GetName()
}
