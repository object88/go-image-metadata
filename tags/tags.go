package tags

import (
	"fmt"

	"github.com/object88/go-image-metadata/common"
)

func defaultInitializer(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	dataSize, ok := common.DataFormatSizes[format]
	if !ok {
		return nil
	}
	switch format {
	case common.ASCIIString:
		return readASCIIString(reader, tag, format, count, data)
	case common.Sbyte, common.Sshort, common.Slong:
		return readSignedInteger(reader, tag, dataSize, format, count, data)
	case common.Srational:
		return readSignedRational(reader, tag, format, count, data)
	case common.Ubyte, common.Ushort, common.Ulong:
		return readUnsignedInteger(reader, tag, dataSize, format, count, data)
	case common.Urational:
		return readUnsignedRational(reader, tag, format, count, data)
	}
	return nil
}

func readASCIIString(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
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
	var v []uint32
	if dataSize*count > 4 {
		v = make([]uint32, count)
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(data))
		// Read off the string of numbers...
		r.SeekTo(cur)
	} else {
		if format == common.Ubyte {
			v, _ = r.ReadUint8FromUint32(count, data)
		} else if format == common.Ushort {
			v, _ = r.ReadUint16FromUint32(count, data)
		} else if format == common.Ulong {
			v = []uint32{data}
		}
	}
	return &UnsignedIntegerMetadata{BaseMetadata{tag}, format, v}
	// fmt.Printf("0x%04x, %s, 0x%08x, 0x%08x\n", tag, format, count, data)
	// return nil
}

func readSignedRational(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(data))
	v := make([]SignedRational, count)
	for i := uint32(0); i < count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = SignedRational{Numerator: int32(n), Denominator: int32(d)}
	}
	r.SeekTo(cur)
	return &SignedRationalMetadata{BaseMetadata{tag}, v}
}

func readUnsignedRational(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(data))
	v := make([]UnsignedRational, count)
	for i := uint32(0); i < count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = UnsignedRational{Numerator: n, Denominator: d}
	}
	r.SeekTo(cur)
	return &UnsignedRationalMetadata{BaseMetadata{tag}, v}
}

// TagMap is the map of all known tags
var TagMap = map[uint16]TagBuilder{
	0x010e: TagBuilder{name: "ImageDescription"},
	0x010f: TagBuilder{name: "Make"},
	0x0110: TagBuilder{name: "Model"},
	0x0112: TagBuilder{name: "Orientation"},
	0x011a: TagBuilder{name: "XResolution"},
	0x011b: TagBuilder{name: "YResolution"},
	0x8769: TagBuilder{
		name: "ExifOffset",
		initializer: func(reader TagReader, tag TagID, format common.DataFormat, count uint32, data uint32) Metadata {
			fmt.Printf("Found ExifBuilder...\n")
			r := reader.GetReader()
			cur := r.GetCurrentOffset()
			reader.ReadIfd(data)
			r.SeekTo(cur)
			return nil
		},
	},
	0x829a: TagBuilder{name: "ExposureTime"},
	0x829d: TagBuilder{name: "FNumber"},
	0x8822: TagBuilder{name: "ExposureProgram"},
	0x8824: TagBuilder{name: "SpectralSensitivity"},
	0x8827: TagBuilder{name: "ISOSpeedRatings"},
	0x8828: TagBuilder{name: "OECF"},
	0x9000: TagBuilder{name: "ExifVersion"},
	0x9003: TagBuilder{name: "DateTimeOriginal"},
	0x9004: TagBuilder{name: "DateTimeDigitized"},
	0x9101: TagBuilder{name: "ComponentsConfiguration"},
	0x9102: TagBuilder{name: "CompressedBitsPerPixel"},
	0x9201: TagBuilder{name: "ShutterSpeedValue"},
	0x9202: TagBuilder{name: "ApertureValue"},
	0x9203: TagBuilder{name: "BrightnessValue"},
	0x9204: TagBuilder{name: "ExposureBiasValue"},
	0x9205: TagBuilder{name: "MaxApertureValue"},
	0x9206: TagBuilder{name: "SubjectDistance"},
	0x9207: TagBuilder{name: "MeteringMode"},
	0x9208: TagBuilder{name: "LightSource"},
	0x9209: TagBuilder{name: "Flash"},
	0x920a: TagBuilder{name: "FocalLength"},
	0x9214: TagBuilder{name: "SubjectArea"},
	0x927c: TagBuilder{name: "MakerNote"},
	0x9286: TagBuilder{name: "UserComment"},
	0x9290: TagBuilder{name: "SubsecTime"},
	0x9291: TagBuilder{name: "SubsecTimeOriginal"},
	0x9292: TagBuilder{name: "SubsecTimeDigitized"},
	0xa000: TagBuilder{name: "FlashpixVersion"},
	0xa001: TagBuilder{name: "ColorSpace"},
	0xa002: TagBuilder{name: "PixelXDimension"},
	0xa003: TagBuilder{name: "PixelYDimension"},
	0xa004: TagBuilder{name: "RelatedSoundFile"},
	0xa20b: TagBuilder{name: "FlashEnergy"},
	0xa20c: TagBuilder{name: "SpatialFrequencyResponse"},
	0xa20e: TagBuilder{name: "FocalPlaneXResolution"},
	0xa20f: TagBuilder{name: "FocalPlaneYResolution"},
	0xa210: TagBuilder{name: "FocalPlaneResolutionUnit"},
	0xa214: TagBuilder{name: "SubjectLocation"},
	0xa215: TagBuilder{name: "ExposureIndex"},
	0xa217: TagBuilder{name: "SensingMethod"},
	0xa300: TagBuilder{name: "FileSource"},
	0xa301: TagBuilder{name: "SceneType"},
	0xa302: TagBuilder{name: "CFAPattern"},
	0xa401: TagBuilder{name: "CustomRendered"},
	0xa402: TagBuilder{name: "ExposureMode"},
	0xa403: TagBuilder{name: "WhiteBalance"},
	0xa404: TagBuilder{name: "DigitalZoomRatio"},
	0xa405: TagBuilder{name: "FocalLengthIn35mmFilm"},
	0xa406: TagBuilder{name: "SceneCaptureType"},
	0xa407: TagBuilder{name: "GainControl"},
	0xa408: TagBuilder{name: "Contrast"},
	0xa409: TagBuilder{name: "Saturation"},
	0xa40a: TagBuilder{name: "Sharpness"},
	0xa40b: TagBuilder{name: "DeviceSettingDescription"},
	0xa40c: TagBuilder{name: "SubjectDistanceRange"},
	0xa420: TagBuilder{name: "ImageUniqueID"},
}
