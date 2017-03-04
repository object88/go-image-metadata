package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// DoubleFloatTag holds an array of 8-byte (64 bit) floats
type DoubleFloatTag struct {
	BaseTag
	value []float64
}

func (m *DoubleFloatTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(common.Dfloat.String())
	buffer.WriteString(") [")
	for k, v := range m.value {
		s32 := strconv.FormatFloat(v, 'E', -1, 64)
		buffer.WriteString(s32)
		if k != len(m.value)-1 {
			buffer.WriteString("\", \"")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
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
