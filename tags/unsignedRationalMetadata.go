package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

type UnsignedRationalMetadata struct {
	BaseMetadata
	value []common.UnsignedRational
}

func (m *UnsignedRationalMetadata) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(common.Urational.String())
	buffer.WriteString(") [")
	for k, v := range m.value {
		buffer.WriteString(strconv.Itoa(int(v.Numerator)))
		buffer.WriteRune('/')
		buffer.WriteString(strconv.Itoa(int(v.Denominator)))
		if k != len(m.value)-1 {
			buffer.WriteString("\", \"")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
