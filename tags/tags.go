package tags

import (
	"fmt"

	"github.com/object88/go-image-metadata/common"
)

type TagBuilder struct {
	Name        string
	Initializer TagInitializer
}

func ReadASCIIString(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	// From the TIFF-v6 spec:
	// Any ASCII field can contain multiple strings, each terminated with a NUL. A
	// single string is preferred whenever possible. The Count for multi-string fields is
	// the number of bytes in all the strings in that field plus their terminating NUL
	// bytes. Only one NUL is allowed between strings, so that the strings following the
	// first string will often begin on an odd byte.
	// ... so this is not sufficient.
	cur := reader.GetReader().GetCurrentOffset()
	reader.GetReader().SeekTo(int64(data))
	s, _ := reader.GetReader().ReadNullTerminatedString()
	reader.GetReader().SeekTo(cur)
	return &StringMetadata{BaseMetadata{tag}, []string{s}}
}

func ReadInteger(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	dataSize, ok := common.DataFormatSizes[format]
	if !ok {
		return nil
	}
	switch format {
	case common.Ubyte, common.Ushort, common.Ulong:
		return readUnsignedInteger(reader, tag, dataSize, format, count, data)
	case common.Sbyte, common.Sshort, common.Slong:
		return readSignedInteger(reader, tag, dataSize, format, count, data)
	default:
		return nil
	}
}

func readSignedInteger(reader TagReader, tag TagID, dataSize uint32, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	v := make([]int32, count)
	if dataSize*count > 4 {
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(data))
		// Read off the string of numbers...
		r.SeekTo(cur)
	} else {
		if format == common.Sbyte {
			for i := uint32(0); i < count; i++ {
				n, _ := r.ReadUint8()
				v[i] = int32(n)
			}
		} else if format == common.Sshort {
			for i := uint32(0); i < count; i++ {
				n, _ := r.ReadUint16()
				v[i] = int32(n)
			}
		} else if format == common.Slong {
			n, _ := r.ReadUint32()
			v[0] = int32(n)
		}
	}
	return &SignedIntegerMetadata{BaseMetadata{tag}, format, v}
}

func readUnsignedInteger(reader TagReader, tag TagID, dataSize uint32, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	v := make([]uint32, count)
	if dataSize*count > 4 {
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(data))
		// Read off the string of numbers...
		r.SeekTo(cur)
	} else {
		if format == common.Ubyte {
			for i := uint32(0); i < count; i++ {
				n, _ := r.ReadUint8()
				v[i] = uint32(n)
			}
		} else if format == common.Ushort {

			for i := uint32(0); i < count; i++ {
				n, _ := r.ReadUint16()
				v[i] = uint32(n)
			}
		} else if format == common.Ulong {
			n, _ := r.ReadUint32()
			v[0] = uint32(n)
		}
	}
	return &UnsignedIntegerMetadata{BaseMetadata{tag}, format, v}
}

func readSignedRational(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(data))
	v := make([]common.SignedRational, count)
	for i := uint32(0); i < count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = common.SignedRational{Numerator: int32(n), Denominator: int32(d)}
	}
	r.SeekTo(cur)
	return &SignedRationalMetadata{BaseMetadata{tag}, v}
}

func readUnsignedRational(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(data))
	v := make([]common.UnsignedRational, count)
	for i := uint32(0); i < count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = common.UnsignedRational{Numerator: n, Denominator: d}
	}
	r.SeekTo(cur)
	return &UnsignedRationalMetadata{BaseMetadata{tag}, v}
}

func ReadRational(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	_, ok := common.DataFormatSizes[format]
	if !ok {
		return nil
	}
	if format == common.Urational {
		return readUnsignedRational(reader, tag, format, count, data)
	} else if format == common.Srational {
		return readSignedRational(reader, tag, format, count, data)
	}
	return nil
}

// TagMap is the map of all known tags
var TagMap = map[uint16]TagBuilder{
	0x010e: TagBuilder{
		Name:        "ImageDescription",
		Initializer: ReadASCIIString,
	},
	0x010f: TagBuilder{
		Name:        "Make",
		Initializer: ReadASCIIString,
	},
	0x0110: TagBuilder{
		Name:        "Model",
		Initializer: ReadASCIIString,
	},
	0x0112: TagBuilder{
		Name:        "Orientation",
		Initializer: ReadInteger,
	},
	0x011a: TagBuilder{
		Name:        "XResolution",
		Initializer: ReadRational,
	},
	0x011b: TagBuilder{
		Name:        "YResolution",
		Initializer: ReadRational,
	},
	0x8769: TagBuilder{
		Name: "ExifOffset",
		Initializer: func(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
			fmt.Printf("Found ExifBuilder...")
			cur := reader.GetReader().GetCurrentOffset()
			reader.ReadIfd(data)
			reader.GetReader().SeekTo(cur)
			return nil
		},
	},
	0x829a: TagBuilder{
		Name:        "ExposureTime",
		Initializer: nil,
	},
	0x829d: TagBuilder{
		Name:        "FNumber",
		Initializer: nil,
	},
	0x8822: TagBuilder{
		Name:        "ExposureProgram",
		Initializer: nil,
	},
	0x8824: TagBuilder{
		Name:        "SpectralSensitivity",
		Initializer: nil,
	},
	0x8827: TagBuilder{
		Name:        "ISOSpeedRatings",
		Initializer: nil,
	},
	0x8828: TagBuilder{
		Name:        "OECF",
		Initializer: nil,
	},
	0x9000: TagBuilder{
		Name:        "ExifVersion",
		Initializer: nil,
	},
	0x9003: TagBuilder{
		Name:        "DateTimeOriginal",
		Initializer: ReadASCIIString,
	},
	0x9004: TagBuilder{
		Name:        "DateTimeDigitized",
		Initializer: ReadASCIIString,
	},
	0x9101: TagBuilder{
		Name:        "ComponentsConfiguration",
		Initializer: nil,
	},
	0x9102: TagBuilder{
		Name:        "CompressedBitsPerPixel",
		Initializer: nil,
	},
	0x9201: TagBuilder{
		Name:        "ShutterSpeedValue",
		Initializer: nil,
	},
	0x9202: TagBuilder{
		Name:        "ApertureValue",
		Initializer: nil,
	},
	0x9203: TagBuilder{
		Name:        "BrightnessValue",
		Initializer: nil,
	},
	0x9204: TagBuilder{
		Name:        "ExposureBiasValue",
		Initializer: nil,
	},
	0x9205: TagBuilder{
		Name:        "MaxApertureValue",
		Initializer: nil,
	},
	0x9206: TagBuilder{
		Name:        "SubjectDistance",
		Initializer: nil,
	},
	0x9207: TagBuilder{
		Name:        "MeteringMode",
		Initializer: nil,
	},
	0x9208: TagBuilder{
		Name:        "LightSource",
		Initializer: nil,
	},
	0x9209: TagBuilder{
		Name:        "Flash",
		Initializer: nil,
	},
	0x920a: TagBuilder{
		Name:        "FocalLength",
		Initializer: nil,
	},
	0x9214: TagBuilder{
		Name:        "SubjectArea",
		Initializer: nil,
	},
	0x927c: TagBuilder{
		Name:        "MakerNote",
		Initializer: nil,
	},
	0x9286: TagBuilder{
		Name:        "UserComment",
		Initializer: nil,
	},
	0x9290: TagBuilder{
		Name:        "SubsecTime",
		Initializer: nil,
	},
	0x9291: TagBuilder{
		Name:        "SubsecTimeOriginal",
		Initializer: nil,
	},
	0x9292: TagBuilder{
		Name:        "SubsecTimeDigitized",
		Initializer: nil,
	},
	0xa000: TagBuilder{
		Name:        "FlashpixVersion",
		Initializer: nil,
	},
	0xa001: TagBuilder{
		Name:        "ColorSpace",
		Initializer: ReadInteger,
	},
	0xa002: TagBuilder{
		Name:        "PixelXDimension",
		Initializer: ReadInteger,
	},
	0xa003: TagBuilder{
		Name:        "PixelYDimension",
		Initializer: ReadInteger,
	},
	0xa004: TagBuilder{
		Name:        "RelatedSoundFile",
		Initializer: nil,
	},
	0xa20b: TagBuilder{
		Name:        "FlashEnergy",
		Initializer: nil,
	},
	0xa20c: TagBuilder{
		Name:        "SpatialFrequencyResponse",
		Initializer: nil,
	},
	0xa20e: TagBuilder{
		Name:        "FocalPlaneXResolution",
		Initializer: nil,
	},
	0xa20f: TagBuilder{
		Name:        "FocalPlaneYResolution",
		Initializer: nil,
	},
	0xa210: TagBuilder{
		Name:        "FocalPlaneResolutionUnit",
		Initializer: ReadInteger,
	},
	0xa214: TagBuilder{
		Name:        "SubjectLocation",
		Initializer: ReadInteger,
	},
	0xa215: TagBuilder{
		Name:        "ExposureIndex",
		Initializer: nil,
	},
	0xa217: TagBuilder{
		Name:        "SensingMethod",
		Initializer: ReadInteger,
	},
	0xa300: TagBuilder{
		Name:        "FileSource",
		Initializer: nil,
	},
	0xa301: TagBuilder{
		Name:        "SceneType",
		Initializer: nil,
	},
	0xa302: TagBuilder{
		Name:        "CFAPattern",
		Initializer: nil,
	},
	0xa401: TagBuilder{
		Name:        "CustomRendered",
		Initializer: nil,
	},
	0xa402: TagBuilder{
		Name:        "ExposureMode",
		Initializer: nil,
	},
	0xa403: TagBuilder{
		Name:        "WhiteBalance",
		Initializer: nil,
	},
	0xa404: TagBuilder{
		Name:        "DigitalZoomRatio",
		Initializer: nil,
	},
	0xa405: TagBuilder{
		Name:        "FocalLengthIn35mmFilm",
		Initializer: nil,
	},
	0xa406: TagBuilder{
		Name:        "SceneCaptureType",
		Initializer: nil,
	},
	0xa407: TagBuilder{
		Name:        "GainControl",
		Initializer: nil,
	},
	0xa408: TagBuilder{
		Name:        "Contrast",
		Initializer: nil,
	},
	0xa409: TagBuilder{
		Name:        "Saturation",
		Initializer: nil,
	},
	0xa40a: TagBuilder{
		Name:        "Sharpness",
		Initializer: nil,
	},
	0xa40b: TagBuilder{
		Name:        "DeviceSettingDescription",
		Initializer: nil,
	},
	0xa40c: TagBuilder{
		Name:        "SubjectDistanceRange",
		Initializer: nil,
	},
	0xa420: TagBuilder{
		Name:        "ImageUniqueID",
		Initializer: nil,
	},
}
