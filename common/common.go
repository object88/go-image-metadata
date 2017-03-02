package common

// ImageReader reads the image tags and stuff.
type ImageReader interface {
	Read() int64
}
