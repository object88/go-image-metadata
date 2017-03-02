package tags

type BaseMetadata struct {
	tagID TagID
}

func (b *BaseMetadata) GetID() TagID {
	return b.tagID
}

func (b *BaseMetadata) GetName() string {
	name := TagMap[uint16(b.tagID)].Name
	return name
}
