package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

type SignedIntegerTag struct {
	BaseTag
	format common.DataFormat
	value  []int32
}

func (m *SignedIntegerTag) String() string {
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
