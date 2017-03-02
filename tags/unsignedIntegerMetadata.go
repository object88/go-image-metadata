package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

type UnsignedIntegerMetadata struct {
	BaseMetadata
	format common.DataFormat
	value  []uint32
}

func (m *UnsignedIntegerMetadata) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(m.format.String())
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
