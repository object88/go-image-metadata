package tags

// TagBuilder contains some initialization information for tags found in
// an image file
type TagBuilder struct {
	name        string
	initializer TagInitializer
}

// GetName returns the tag name related to this builder
func (tb *TagBuilder) GetName() string {
	return tb.name
}

// GetInitializer returns the tag initializer function.  If the initializer
// was not specified, a default initializer will be used.
func (tb *TagBuilder) GetInitializer() TagInitializer {
	if tb.initializer != nil {
		return tb.initializer
	}

	return defaultInitializer
}
