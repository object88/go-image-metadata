package tags

import "bytes"

type StringTag struct {
	BaseTag
	value []string
}

func (m *StringTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" [\"")
	for k, v := range m.value {
		buffer.WriteString(v)
		if k != len(m.value)-1 {
			buffer.WriteString("\", \"")
		}
	}
	buffer.WriteString("\"]")
	return buffer.String()
}
