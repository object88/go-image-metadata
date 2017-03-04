package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// SingleFloatTag holds an array of 4-byte (32 bit) floats
type SingleFloatTag struct {
	BaseTag
	value []float32
}

func (m *SingleFloatTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(common.Sfloat.String())
	buffer.WriteString(") [")
	for k, v := range m.value {
		s32 := strconv.FormatFloat(float64(v), 'E', -1, 32)
		buffer.WriteString(s32)
		if k != len(m.value)-1 {
			buffer.WriteString("\", \"")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func readSingleFloat(reader TagReader, name string, raw *RawTagData) (Tag, bool, error) {
	r := reader.GetReader()
	v := make([]float32, raw.Count)
	if raw.Count == 1 {
		n, _ := r.ReadUint32()
		v[0] = float32(n)
	} else {
		cur := r.GetCurrentOffset()
		r.SeekTo(int64(raw.Data))
		for i := uint32(0); i < raw.Count; i++ {
			n, _ := r.ReadUint32()
			v[i] = float32(n)
		}
		r.SeekTo(cur)
	}
	return &SingleFloatTag{BaseTag{name, raw.Tag, raw.Format}, v}, true, nil
}
