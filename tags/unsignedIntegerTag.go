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
	value []uint32
}

func (m *UnsignedIntegerTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(m.BaseTag.GetType().String())
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

func readUnsignedInteger(reader TagReader, name string, dataSize uint32, raw *RawTagData) (Tag, bool, error) {
	r := reader.GetReader()
	var v []uint32
	if dataSize*raw.Count > 4 {
		v = make([]uint32, raw.Count)
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(raw.Data))
		// Read off the string of numbers...
		r.SeekTo(cur)
	} else {
		if raw.Format == common.Ubyte {
			v, _ = r.ReadUint8FromUint32(raw.Count, raw.Data)
		} else if raw.Format == common.Ushort {
			v, _ = r.ReadUint16FromUint32(raw.Count, raw.Data)
		} else if raw.Format == common.Ulong {
			v = []uint32{raw.Data}
		}
	}
	return &UnsignedIntegerTag{BaseTag{name, raw.Tag, raw.Format}, v}, true, nil
}
