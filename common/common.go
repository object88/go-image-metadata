package common

// ImageReader reads the image metadata and stuff.
type ImageReader interface {
	Read() int64
}

// SignedRational is a fractional representation of an signed number
type SignedRational struct {
	Numerator   int32
	Denominator int32
}

// UnsignedRational is a fractional representation of an unsigned number
type UnsignedRational struct {
	Numerator   uint32
	Denominator uint32
}
