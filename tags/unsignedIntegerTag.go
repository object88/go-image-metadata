package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// UnsignedIntegerTag holds an array of integers.  All values are stored as
// 32 bits, but the type represents 8, 16, and 32 bit unsigned integers.
type UnsignedIntegerTag struct {
	BaseTag
	format common.DataFormat
	value  []uint32
}

func (m *UnsignedIntegerTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(m.format.String())
	buffer.WriteString(") [")
	for k, v := range m.value {
		buffer.WriteString(strconv.Itoa(int(v)))
		if k != len(m.value)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
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
