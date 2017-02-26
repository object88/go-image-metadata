package tiff

type dataFormat uint16

const (
	_ dataFormat = iota
	ubyte
	asciiString
	ushort
	ulong
	urational
	sbyte
	_
	sshort
	slong
	srational
	sfloat
	dfloat
)

type dataFormatMetadata struct {
	size uint8
}

var dataFormatMetadatas = map[dataFormat]dataFormatMetadata{
	ubyte: dataFormatMetadata{
		size: 1,
	},
	asciiString: dataFormatMetadata{
		size: 1,
	},
}

var dataFormats = [...]string{
	"",
	"unsignd byte",
	"ascii string",
	"unsigned short",
	"unsigned long",
	"unsigned rational",
	"signed byte",
	"",
	"signed short",
	"signed long",
	"signed rational",
	"single float",
	"double float",
}

func (df dataFormat) String() string {
	return dataFormats[df]
}
