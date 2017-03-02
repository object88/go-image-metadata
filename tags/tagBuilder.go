package tags

type TagBuilder struct {
	name        string
	initializer TagInitializer
}

func (tb *TagBuilder) GetName() string {
	return tb.name
}

func (tb *TagBuilder) GetInitializer() TagInitializer {
	if tb.initializer != nil {
		return tb.initializer
	}

	return defaultInitializer
}
