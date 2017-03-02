package tags

type BaseTag struct {
	tagID TagID
}

func (b *BaseTag) GetID() TagID {
	return b.tagID
}

func (b *BaseTag) GetName() string {
	tagBuilder := TagMap[uint16(b.tagID)]
	return tagBuilder.GetName()
}
