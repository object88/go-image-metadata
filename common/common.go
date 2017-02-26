package common

// ImageReader reads the image metadata and stuff.
type ImageReader interface {
	Read() int64
}
