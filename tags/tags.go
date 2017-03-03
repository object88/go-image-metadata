package tags

import (
	"errors"
	"fmt"

	"github.com/object88/go-image-metadata/common"
)

func defaultInitializer(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
	dataSize, ok := common.DataFormatSizes[format]
	if !ok {
		return nil, false, errors.New("Do not have matching data format size")
	}
	switch format {
	case common.ASCIIString:
		return readASCIIString(reader, tag, name, format, count, data)
	case common.Dfloat:
		return readDoubleFloat(reader, tag, name, format, count, data)
	case common.Sbyte, common.Sshort, common.Slong:
		return readSignedInteger(reader, tag, name, dataSize, format, count, data)
	case common.Sfloat:
		return readSingleFloat(reader, tag, name, format, count, data)
	case common.Srational:
		return readSignedRational(reader, tag, name, format, count, data)
	case common.Ubyte, common.Ushort, common.Ulong:
		return readUnsignedInteger(reader, tag, name, dataSize, format, count, data)
	case common.Urational:
		return readUnsignedRational(reader, tag, name, format, count, data)
	}
	return nil, false, nil
}

func readASCIIString(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
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
	return &StringTag{BaseTag{name, tag}, []string{s}}, true, nil
}

func readDoubleFloat(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(data))
	v := make([]float64, count)
	for i := uint32(0); i < count; i++ {
		n, _ := r.ReadUint64()
		v[i] = float64(n)
	}
	r.SeekTo(cur)
	return &DoubleFloatTag{BaseTag{name, tag}, v}, true, nil
}

func readSignedInteger(reader TagReader, tag TagID, name string, dataSize uint32, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
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
	return &SignedIntegerTag{BaseTag{name, tag}, format, v}, true, nil
}

func readSignedRational(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
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
	return &SignedRationalTag{BaseTag{name, tag}, v}, true, nil
}

func readSingleFloat(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
	r := reader.GetReader()
	v := make([]float32, count)
	if count == 1 {
		n, _ := r.ReadUint32()
		v[0] = float32(n)
	} else {
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(data))
		for i := uint32(0); i < count; i++ {
			n, _ := r.ReadUint32()
			v[i] = float32(n)
		}
		r.SeekTo(cur)
	}
	return &SingleFloatTag{BaseTag{name, tag}, v}, true, nil
}

func readUnsignedInteger(reader TagReader, tag TagID, name string, dataSize uint32, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
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
	return &UnsignedIntegerTag{BaseTag{name, tag}, format, v}, true, nil
}

func readUnsignedRational(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
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
	return &UnsignedRationalTag{BaseTag{name, tag}, v}, true, nil
}

// TagMap is the map of all known tags
var TagMap = map[uint16]TagBuilder{
	0x0103: TagBuilder{name: "Compression"},
	0x010e: TagBuilder{name: "ImageDescription"},
	0x010f: TagBuilder{name: "Make"},
	0x0110: TagBuilder{name: "Model"},
	0x0112: TagBuilder{name: "Orientation"},
	0x011a: TagBuilder{name: "XResolution"},
	0x011b: TagBuilder{name: "YResolution"},
	0x0128: TagBuilder{name: "ResolutionUnit"},
	0x0131: TagBuilder{name: "Software"},
	0x0132: TagBuilder{name: "DateTime"},
	0x0201: TagBuilder{name: "JPEGInterchangeFormat"},
	0x0202: TagBuilder{name: "JPEGInterchangeFormatLength"},
	0x8825: TagBuilder{
		name: "GPS IFD",
		initializer: func(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
			fmt.Printf("Found GPS IFD...\n")
			r := reader.GetReader()
			cur := r.GetCurrentOffset()
			reader.ReadIfd(data, GpsTagMap)
			r.SeekTo(cur)
			return nil, true, nil
		},
	},
	0x8298: TagBuilder{name: "Copyright"},
	0x8769: TagBuilder{
		name: "ExifOffset",
		initializer: func(reader TagReader, tag TagID, name string, format common.DataFormat, count uint32, data uint32) (Tag, bool, error) {
			fmt.Printf("Found ExifBuilder...\n")
			r := reader.GetReader()
			cur := r.GetCurrentOffset()
			reader.ReadIfd(data, ExifTagMap)
			r.SeekTo(cur)
			return nil, true, nil
		},
	},
}

// ExifTagMap contains the tags related to Exit data.
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/privateifd/exif.html
var ExifTagMap = map[uint16]TagBuilder{
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

// GpsTagMap contains the tags related to GPS data.
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/privateifd/gps.html
var GpsTagMap = map[uint16]TagBuilder{
	0x0000: TagBuilder{name: "GPSVersionID"},
	0x0001: TagBuilder{name: "GPSLatitudeRef"},
	0x0002: TagBuilder{name: "GPSLatitude"},
	0x0003: TagBuilder{name: "GPSLongitudeRef"},
	0x0004: TagBuilder{name: "GPSLongitude"},
	0x0005: TagBuilder{name: "GPSAltitudeRef"},
	0x0006: TagBuilder{name: "GPSAltitude"},
	0x0007: TagBuilder{name: "GPSTimeStamp"},
	0x0008: TagBuilder{name: "GPSSatellites"},
	0x0009: TagBuilder{name: "GPSStatus"},
	0x000A: TagBuilder{name: "GPSMeasureMode"},
	0x000B: TagBuilder{name: "GPSDOP"},
	0x000C: TagBuilder{name: "GPSSpeedRef"},
	0x000D: TagBuilder{name: "GPSSpeed"},
	0x000E: TagBuilder{name: "GPSTrackRef"},
	0x000F: TagBuilder{name: "GPSTrack"},
	0x0010: TagBuilder{name: "GPSImgDirectionRef"},
	0x0011: TagBuilder{name: "GPSImgDirection"},
	0x0012: TagBuilder{name: "GPSMapDatum"},
	0x0013: TagBuilder{name: "GPSDestLatitudeRef"},
	0x0014: TagBuilder{name: "GPSDestLatitude"},
	0x0015: TagBuilder{name: "GPSDestLongitudeRef"},
	0x0016: TagBuilder{name: "GPSDestLongitude"},
	0x0017: TagBuilder{name: "GPSDestBearingRef"},
	0x0018: TagBuilder{name: "GPSDestBearing"},
	0x0019: TagBuilder{name: "GPSDestDistanceRef"},
	0x001A: TagBuilder{name: "GPSDestDistance"},
	0x001B: TagBuilder{name: "GPSProcessingMethod"},
	0x001C: TagBuilder{name: "GPSAreaInformation"},
	0x001D: TagBuilder{name: "GPSDateStamp"},
	0x001E: TagBuilder{name: "GPSDifferential"},
}
