package common

// TagID is a metadata identifier
type TagID uint16

type TagStruct struct {
	ID TagID
}

const (
	// ImageDescription describes the image
	ImageDescription TagID = 0x010e

	// Make is the make of the camera
	Make = 0x010f

	// Model is the model of the camera
	Model = 0x0110

	// ExifOffset points to a memory location for EXIF data.  Internal use only.
	ExifOffset = 0x8769
)
