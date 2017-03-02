package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

type DoubleFloatMetadata struct {
	BaseMetadata
	value []float64
}

func (m *DoubleFloatMetadata) String() string {
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
