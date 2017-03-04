package tags

import (
	"errors"
	"fmt"

	"github.com/object88/go-image-metadata/common"
)

// TagMap is the map of baseline and extension tags
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/baseline.html
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/extension.html
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/private.html
var TagMap map[uint16]TagBuilder

// ExifTagMap contains the tags related to Exit data.
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/privateifd/exif.html
var ExifTagMap map[uint16]TagBuilder

// GpsTagMap contains the tags related to GPS data.
// Ref: http://www.awaresystems.be/imaging/tiff/tifftags/privateifd/gps.html
var GpsTagMap map[uint16]TagBuilder

// InteropTagMap contains the tags related to interoperability data.
// REf: http://www.awaresystems.be/imaging/tiff/tifftags/privateifd/interoperability.html
var InteropTagMap map[uint16]TagBuilder

func init() {
	TagMap = map[uint16]TagBuilder{
		0x00fe: TagBuilder{name: "NewSubfileType"},
		0x00ff: TagBuilder{name: "SubfileType"},
		0x0100: TagBuilder{name: "ImageWidth"},
		0x0101: TagBuilder{name: "ImageLength"},
		0x0102: TagBuilder{name: "BitsPerSample"},
		0x0103: TagBuilder{name: "Compression"},
		0x0106: TagBuilder{name: "PhotometricInterpretation"},
		0x0107: TagBuilder{name: "Threshholding"},
		0x0108: TagBuilder{name: "CellWidth"},
		0x0109: TagBuilder{name: "CellLength"},
		0x010a: TagBuilder{name: "FillOrder"},
		0x010e: TagBuilder{name: "ImageDescription"},
		0x010f: TagBuilder{name: "Make"},
		0x0110: TagBuilder{name: "Model"},
		0x0111: TagBuilder{name: "StripOffsets"},
		0x0112: TagBuilder{name: "Orientation"},
		0x0115: TagBuilder{name: "SamplesPerPixel"},
		0x0116: TagBuilder{name: "RowsPerStrip"},
		0x0117: TagBuilder{name: "StripByteCounts"},
		0x0118: TagBuilder{name: "MinSampleValue"},
		0x0119: TagBuilder{name: "MaxSampleValue"},
		0x011a: TagBuilder{name: "XResolution"},
		0x011b: TagBuilder{name: "YResolution"},
		0x011c: TagBuilder{name: "PlanarConfiguration"},
		0x0120: TagBuilder{name: "FreeOffsets"},
		0x0121: TagBuilder{name: "FreeByteCounts"},
		0x0122: TagBuilder{name: "GrayResponseUnit"},
		0x0123: TagBuilder{name: "GrayResponseCurve"},
		0x0128: TagBuilder{name: "ResolutionUnit"},
		0x0131: TagBuilder{name: "Software"},
		0x0132: TagBuilder{name: "DateTime"},
		0x013b: TagBuilder{name: "Artist"},
		0x013c: TagBuilder{name: "HostComputer"},
		0x0140: TagBuilder{name: "ColorMap"},
		0x0152: TagBuilder{name: "ExtraSamples"},
		0x8298: TagBuilder{name: "Copyright"},

		0x010d: TagBuilder{name: "DocumentName"},
		0x011d: TagBuilder{name: "PageName"},
		0x011e: TagBuilder{name: "XPosition"},
		0x011f: TagBuilder{name: "YPosition"},
		0x0124: TagBuilder{name: "T4Options"},
		0x0125: TagBuilder{name: "T6Options"},
		0x0129: TagBuilder{name: "PageNumber"},
		0x012d: TagBuilder{name: "TransferFunction"},
		0x013d: TagBuilder{name: "Predictor"},
		0x013e: TagBuilder{name: "WhitePoint"},
		0x013f: TagBuilder{name: "PrimaryChromaticities"},
		0x0141: TagBuilder{name: "HalftoneHints"},
		0x0142: TagBuilder{name: "TileWidth"},
		0x0143: TagBuilder{name: "TileLength"},
		0x0144: TagBuilder{name: "TileOffsets"},
		0x0145: TagBuilder{name: "TileByteCounts"},
		0x0146: TagBuilder{name: "BadFaxLines"},
		0x0147: TagBuilder{name: "CleanFaxData"},
		0x0148: TagBuilder{name: "ConsecutiveBadFaxLines"},
		0x014a: TagBuilder{name: "SubIFDs"},
		0x014c: TagBuilder{name: "InkSet"},
		0x014d: TagBuilder{name: "InkNames"},
		0x014e: TagBuilder{name: "NumberOfInks"},
		0x0150: TagBuilder{name: "DotRange"},
		0x0151: TagBuilder{name: "TargetPrinter"},
		0x0153: TagBuilder{name: "SampleFormat"},
		0x0154: TagBuilder{name: "SMinSampleValue"},
		0x0155: TagBuilder{name: "SMaxSampleValue"},
		0x0156: TagBuilder{name: "TransferRange"},
		0x0157: TagBuilder{name: "ClipPath"},
		0x0158: TagBuilder{name: "XClipPathUnits"},
		0x0159: TagBuilder{name: "YClipPathUnits"},
		0x015A: TagBuilder{name: "Indexed"},
		0x015B: TagBuilder{name: "JPEGTables"},
		0x015F: TagBuilder{name: "OPIProxy"},
		0x0190: TagBuilder{name: "GlobalParametersIFD"},
		0x0191: TagBuilder{name: "ProfileType"},
		0x0192: TagBuilder{name: "FaxProfile"},
		0x0193: TagBuilder{name: "CodingMethods"},
		0x0194: TagBuilder{name: "VersionYear"},
		0x0195: TagBuilder{name: "ModeNumber"},
		0x01b1: TagBuilder{name: "Decode"},
		0x01b2: TagBuilder{name: "DefaultImageColor"},
		0x0200: TagBuilder{name: "JPEGProc"},
		0x0201: TagBuilder{name: "JPEGInterchangeFormat"},
		0x0202: TagBuilder{name: "JPEGInterchangeFormatLength"},
		0x0203: TagBuilder{name: "JPEGRestartInterval"},
		0x0205: TagBuilder{name: "JPEGLosslessPredictors"},
		0x0206: TagBuilder{name: "JPEGPointTransforms"},
		0x0207: TagBuilder{name: "JPEGQTables"},
		0x0208: TagBuilder{name: "JPEGDCTables"},
		0x0209: TagBuilder{name: "JPEGACTables"},
		0x0211: TagBuilder{name: "YCbCrCoefficients"},
		0x0212: TagBuilder{name: "YCbCrSubSampling"},
		0x0213: TagBuilder{name: "YCbCrPositioning"},
		0x0214: TagBuilder{name: "ReferenceBlackWhite"},
		0x022f: TagBuilder{name: "StripRowCounts"},
		0x02bc: TagBuilder{name: "XMP"},
		0x800d: TagBuilder{name: "ImageID"},
		0x87ac: TagBuilder{name: "ImageLayer"},

		0x80a4: TagBuilder{name: "Wang Annotation"},
		0x82a5: TagBuilder{name: "MD FileTAg"},
		0x82a6: TagBuilder{name: "MD ScalePixel"},
		0x82a7: TagBuilder{name: "MD ColorTable"},
		0x82a8: TagBuilder{name: "MD LabName"},
		0x82a9: TagBuilder{name: "MD SampleInfo"},
		0x82aa: TagBuilder{name: "MD PrepDate"},
		0x82ab: TagBuilder{name: "MD PrepTime"},
		0x82ac: TagBuilder{name: "MD FileUnits"},
		0x830e: TagBuilder{name: "ModelPixelScaleTag"},
		0x83bb: TagBuilder{name: "IPTC"},
		0x847e: TagBuilder{name: "INGR Packet Data Tag"},
		0x847f: TagBuilder{name: "INGR Flag Registers"},
		0x8480: TagBuilder{name: "IrasB Transformation Matrix"},
		0x8482: TagBuilder{name: "ModelTiepointTag"},
		0x85d8: TagBuilder{name: "ModelTransformationTag"},
		0x8649: TagBuilder{name: "Photoshop"},
		0x8769: TagBuilder{
			name: "Exif IFD",
			initializer: func(reader TagReader, foundTags *map[uint16]Tag, name string, raw *RawTagData) (Tag, bool, error) {
				r := reader.GetReader()
				cur := r.GetCurrentOffset()
				reader.ReadIfd(raw.Data, []*map[uint16]TagBuilder{&ExifTagMap, &TagMap}, foundTags)
				r.SeekTo(cur)
				return nil, true, nil
			},
		},
		0x8773: TagBuilder{name: "ICC Profile"},
		0x87af: TagBuilder{name: "GeoKeyDirectoryTag"},
		0x87B0: TagBuilder{name: "GeoDoubleParamsTag"},
		0x87B1: TagBuilder{name: "GeoAsciiParamsTag"},
		0x8825: TagBuilder{
			name: "GPS IFD",
			initializer: func(reader TagReader, foundTags *map[uint16]Tag, name string, raw *RawTagData) (Tag, bool, error) {
				fmt.Printf("Found GPS IFD...\n")
				r := reader.GetReader()
				cur := r.GetCurrentOffset()
				reader.ReadIfd(raw.Data, []*map[uint16]TagBuilder{&GpsTagMap}, foundTags)
				r.SeekTo(cur)
				return nil, true, nil
			},
		},
		0x885C: TagBuilder{name: "HylaFAX FaxRecvParams"},
		0x885D: TagBuilder{name: "HylaFAX FaxSubAddress"},
		0x885E: TagBuilder{name: "HylaFAX FaxRecvTime"},
		0x935C: TagBuilder{name: "ImageSourceData"},
		0xa005: TagBuilder{name: "Interoperability"},
		0xa480: TagBuilder{name: "GDAL_METADATA"},
		0xa481: TagBuilder{name: "GDAL_NODATA"},
		0xc427: TagBuilder{name: "Oce Scanjob Description"},
		0xc428: TagBuilder{name: "Oce Application Selector"},
		0xc429: TagBuilder{name: "Oce Identification Number"},
		0xc42A: TagBuilder{name: "Oce ImageLogic Characteristics"},
		0xc612: TagBuilder{name: "DNGVersion"},
		0xc613: TagBuilder{name: "DNGBackwardVersion"},
		0xc614: TagBuilder{name: "UniqueCameraModel"},
		0xc615: TagBuilder{name: "LocalizedCameraModel"},
		0xc616: TagBuilder{name: "CFAPlaneColor"},
		0xc617: TagBuilder{name: "CFALayout"},
		0xc618: TagBuilder{name: "LinearizationTable"},
		0xc619: TagBuilder{name: "BlackLevelRepeatDim"},
		0xc61A: TagBuilder{name: "BlackLevel"},
		0xc61B: TagBuilder{name: "BlackLevelDeltaH"},
		0xc61C: TagBuilder{name: "BlackLevelDeltaV"},
		0xc61D: TagBuilder{name: "WhiteLevel"},
		0xc61E: TagBuilder{name: "DefaultScale"},
		0xc61F: TagBuilder{name: "DefaultCropOrigin"},
		0xc620: TagBuilder{name: "DefaultCropSize"},
		0xc621: TagBuilder{name: "ColorMatrix1"},
		0xc622: TagBuilder{name: "ColorMatrix2"},
		0xc623: TagBuilder{name: "CameraCalibration1"},
		0xc624: TagBuilder{name: "CameraCalibration2"},
		0xc625: TagBuilder{name: "ReductionMatrix1"},
		0xc626: TagBuilder{name: "ReductionMatrix2"},
		0xc627: TagBuilder{name: "AnalogBalance"},
		0xc628: TagBuilder{name: "AsShotNeutral"},
		0xc629: TagBuilder{name: "AsShotWhiteXY"},
		0xc62A: TagBuilder{name: "BaselineExposure"},
		0xc62B: TagBuilder{name: "BaselineNoise"},
		0xc62C: TagBuilder{name: "BaselineSharpness"},
		0xc62D: TagBuilder{name: "BayerGreenSplit"},
		0xc62E: TagBuilder{name: "LinearResponseLimit"},
		0xc62F: TagBuilder{name: "CameraSerialNumber"},
		0xc630: TagBuilder{name: "LensInfo"},
		0xc631: TagBuilder{name: "ChromaBlurRadius"},
		0xc632: TagBuilder{name: "AntiAliasStrength"},
		0xc634: TagBuilder{name: "DNGPrivateData"},
		0xc635: TagBuilder{name: "MakerNoteSafety"},
		0xc65A: TagBuilder{name: "CalibrationIlluminant1"},
		0xc65B: TagBuilder{name: "CalibrationIlluminant2"},
		0xc65C: TagBuilder{name: "BestQualityScale"},
		0xc660: TagBuilder{name: "Alias Layer Metadata"},
	}

	ExifTagMap = map[uint16]TagBuilder{
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

	GpsTagMap = map[uint16]TagBuilder{
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
		0x000a: TagBuilder{name: "GPSMeasureMode"},
		0x000b: TagBuilder{name: "GPSDOP"},
		0x000c: TagBuilder{name: "GPSSpeedRef"},
		0x000d: TagBuilder{name: "GPSSpeed"},
		0x000e: TagBuilder{name: "GPSTrackRef"},
		0x000f: TagBuilder{name: "GPSTrack"},
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
		0x001a: TagBuilder{name: "GPSDestDistance"},
		0x001b: TagBuilder{name: "GPSProcessingMethod"},
		0x001c: TagBuilder{name: "GPSAreaInformation"},
		0x001d: TagBuilder{name: "GPSDateStamp"},
		0x001e: TagBuilder{name: "GPSDifferential"},
	}

	InteropTagMap = map[uint16]TagBuilder{
		0x0001: TagBuilder{name: "InteroperabilityIndex"},
	}
}

func defaultInitializer(reader TagReader, _ *map[uint16]Tag, name string, raw *RawTagData) (Tag, bool, error) {
	dataSize, ok := common.DataFormatSizes[raw.Format]
	if !ok {
		return nil, false, errors.New("Do not have matching data format size")
	}
	switch raw.Format {
	case common.ASCIIString:
		return readASCIIString(reader, name, raw)
	case common.Dfloat:
		return readDoubleFloat(reader, name, raw)
	case common.Sbyte, common.Sshort, common.Slong:
		return readSignedInteger(reader, name, dataSize, raw)
	case common.Sfloat:
		return readSingleFloat(reader, name, raw)
	case common.Srational:
		return readSignedRational(reader, name, raw)
	case common.Ubyte, common.Ushort, common.Ulong:
		return readUnsignedInteger(reader, name, dataSize, raw)
	case common.Urational:
		return readUnsignedRational(reader, name, raw)
	}
	return nil, false, nil
}
