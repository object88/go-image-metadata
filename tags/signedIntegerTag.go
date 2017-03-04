package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// SignedIntegerTag holds an array of integers.  All values are stored as
// 32 bits, but the type represents 8, 16, and 32 bit signed integers.
type SignedIntegerTag struct {
	BaseTag
	value []int32
}

func (m *SignedIntegerTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(m.BaseTag.GetType().String())
	buffer.WriteString(") [")
	for k, v := range m.value {
		buffer.WriteString(strconv.Itoa(int(v)))
		if k != len(m.value)-1 {
			buffer.WriteString("\", \"")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func readSignedInteger(reader TagReader, name string, dataSize uint32, raw *RawTagData) (Tag, bool, error) {
	r := reader.GetReader()
	v := make([]int32, raw.Count)
	if dataSize*raw.Count > 4 {
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(raw.Data))
		// Read off the string of numbers...
		r.SeekTo(cur)
	} else {
		if raw.Format == common.Sbyte {
			for i := uint32(0); i < raw.Count; i++ {
				n, _ := r.ReadUint8()
				v[i] = int32(n)
			}
		} else if raw.Format == common.Sshort {
			for i := uint32(0); i < raw.Count; i++ {
				n, _ := r.ReadUint16()
				v[i] = int32(n)
			}
		} else if raw.Format == common.Slong {
			n, _ := r.ReadUint32()
			v[0] = int32(n)
		}
	}
	return &SignedIntegerTag{BaseTag{name, raw.Tag, raw.Format}, v}, true, nil
}
