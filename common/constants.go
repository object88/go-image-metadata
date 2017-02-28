package common

// TagID is a metadata identifier
type TagID uint16

type TagStruct struct {
	Name string
}

const (
	ExifOffset uint16 = 0x8769
)

// TagMap is the map of all known tags
var TagMap = map[uint16]TagStruct{
	0x010e: TagStruct{
		Name: "ImageDescription",
	},
	0x010f: TagStruct{
		Name: "Make",
	},
	0x0110: TagStruct{
		Name: "Model",
	},
	0x8769: TagStruct{
		Name: "ExifOffset",
	},
	0x829a: TagStruct{
		Name: "ExposureTime",
	},
	0x829d: TagStruct{
		Name: "FNumber",
	},
	0x8822: TagStruct{
		Name: "ExposureProgram",
	},
	0x8824: TagStruct{
		Name: "SpectralSensitivity",
	},
	0x8827: TagStruct{
		Name: "ISOSpeedRatings",
	},
	0x8828: TagStruct{
		Name: "OECF",
	},
	0x9000: TagStruct{
		Name: "ExifVersion",
	},
	0x9003: TagStruct{
		Name: "DateTimeOriginal",
	},
	0x9004: TagStruct{
		Name: "DateTimeDigitized",
	},
	0x9101: TagStruct{
		Name: "ComponentsConfiguration",
	},
	0x9102: TagStruct{
		Name: "CompressedBitsPerPixel",
	},
	0x9201: TagStruct{
		Name: "ShutterSpeedValue",
	},
	0x9202: TagStruct{
		Name: "ApertureValue",
	},
	0x9203: TagStruct{
		Name: "BrightnessValue",
	},
	0x9204: TagStruct{
		Name: "ExposureBiasValue",
	},
	0x9205: TagStruct{
		Name: "MaxApertureValue",
	},
	0x9206: TagStruct{
		Name: "SubjectDistance",
	},
	0x9207: TagStruct{
		Name: "MeteringMode",
	},
	0x9208: TagStruct{
		Name: "LightSource",
	},
	0x9209: TagStruct{
		Name: "Flash",
	},
	0x920a: TagStruct{
		Name: "FocalLength",
	},
	0x9214: TagStruct{
		Name: "SubjectArea",
	},
	0x927c: TagStruct{
		Name: "MakerNote",
	},
	0x9286: TagStruct{
		Name: "UserComment",
	},
	0x9290: TagStruct{
		Name: "SubsecTime",
	},
	0x9291: TagStruct{
		Name: "SubsecTimeOriginal",
	},
	0x9292: TagStruct{
		Name: "SubsecTimeDigitized",
	},
	0xa000: TagStruct{
		Name: "FlashpixVersion",
	},
	0xa001: TagStruct{
		Name: "ColorSpace",
	},
	0xa002: TagStruct{
		Name: "PixelXDimension",
	},
	0xa003: TagStruct{
		Name: "PixelYDimension",
	},
	0xa004: TagStruct{
		Name: "RelatedSoundFile",
	},
	0xa20b: TagStruct{
		Name: "FlashEnergy",
	},
	0xa20c: TagStruct{
		Name: "SpatialFrequencyResponse",
	},
	0xa20e: TagStruct{
		Name: "FocalPlaneXResolution",
	},
	0xa20f: TagStruct{
		Name: "FocalPlaneYResolution",
	},
	0xa210: TagStruct{
		Name: "FocalPlaneResolutionUnit",
	},
	0xa214: TagStruct{
		Name: "SubjectLocation",
	},
	0xa215: TagStruct{
		Name: "ExposureIndex",
	},
	0xa217: TagStruct{
		Name: "SensingMethod",
	},
	0xa300: TagStruct{
		Name: "FileSource",
	},
	0xa301: TagStruct{
		Name: "SceneType",
	},
	0xa302: TagStruct{
		Name: "CFAPattern",
	},
	0xa401: TagStruct{
		Name: "CustomRendered",
	},
	0xa402: TagStruct{
		Name: "ExposureMode",
	},
	0xa403: TagStruct{
		Name: "WhiteBalance",
	},
	0xa404: TagStruct{
		Name: "DigitalZoomRatio",
	},
	0xa405: TagStruct{
		Name: "FocalLengthIn35mmFilm",
	},
	0xa406: TagStruct{
		Name: "SceneCaptureType",
	},
	0xa407: TagStruct{
		Name: "GainControl",
	},
	0xa408: TagStruct{
		Name: "Contrast",
	},
	0xa409: TagStruct{
		Name: "Saturation",
	},
	0xa40a: TagStruct{
		Name: "Sharpness",
	},
	0xa40b: TagStruct{
		Name: "DeviceSettingDescription",
	},
	0xa40c: TagStruct{
		Name: "SubjectDistanceRange",
	},
	0xa420: TagStruct{
		Name: "ImageUniqueID",
	},
}
