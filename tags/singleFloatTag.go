package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

type SingleFloatMetadata struct {
	BaseMetadata
	value []float32
}

func (m *SingleFloatMetadata) String() string {
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
