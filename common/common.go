package common

// Metadata is a single name-value pair.
type Metadata struct {
	name string
}

type IntegerMetadata struct {
	Metadata
	value int64
}

type StringMetadata struct {
	Metadata
	value string
}

// ImageReader reads the image metadata and stuff.
type ImageReader interface {
	Read() int64
}
