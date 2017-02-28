package common

// DataFormat is the encoding of a piece of data, as described by the TIFF
// standard
type DataFormat uint16

const (
	_ DataFormat = iota

	// Ubyte is an unsigned byte
	Ubyte

	// ASCIIString is a string of 7-bit characters
	ASCIIString

	// Ushort is an unsigned short (2 bytes)
	Ushort

	// Ulong is an unsigned long (4 bytes)
	Ulong

	// Urational is an 8-byte structure; two sequential unsigned longs,
	// representing a numerator and demoninator
	Urational

	// Sbyte is a signed byte
	Sbyte

	// Undefined is special.
	Undefined

	// Sshort is a signed short (2 bytes)
	Sshort

	// Slong is a signed long (4 bytes)
	Slong

	// Srational is an 8-byte structure; two sequential signed longs,
	// representing a numerator and demoninator
	Srational

	// Sfloat is a signed 4 byte floating point number
	Sfloat

	// Dfloat is a signed 8 byte floating point number
	Dfloat
)

// DataFormatSizes maps a DataFormat to the number of bytes a single instance
// requires.
var DataFormatSizes = map[DataFormat]uint32{
	Ubyte:       1,
	ASCIIString: 1,
	Ushort:      2,
	Ulong:       4,
	Urational:   8,
	Undefined:   1,
	Sbyte:       1,
	Sshort:      2,
	Slong:       4,
	Srational:   8,
	Sfloat:      4,
	Dfloat:      8,
}

var dataFormats = [...]string{
	"",
	"unsignd byte",
	"ascii string",
	"unsigned short",
	"unsigned long",
	"unsigned rational",
	"signed byte",
	"undefined",
	"signed short",
	"signed long",
	"signed rational",
	"single float",
	"double float",
}

func (df DataFormat) String() string {
	return dataFormats[df]
}
