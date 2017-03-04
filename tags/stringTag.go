package tags

import (
	"bytes"
)

// StringTag holds an array of strings
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

func readASCIIString(reader TagReader, name string, raw *RawTagData) (Tag, bool, error) {
	// From the TIFF-v6 spec:
	// Any ASCII field can contain multiple strings, each terminated with a NUL. A
	// single string is preferred whenever possible. The Count for multi-string fields is
	// the number of bytes in all the strings in that field plus their terminating NUL
	// bytes. Only one NUL is allowed between strings, so that the strings following the
	// first string will often begin on an odd byte.
	// ... so this is not sufficient.
	cur := reader.GetReader().GetCurrentOffset()
	reader.GetReader().SeekTo(int64(raw.Data))
	s, _ := reader.GetReader().ReadNullTerminatedString()
	reader.GetReader().SeekTo(cur)
	return &StringTag{BaseTag{name, raw.Tag, raw.Format}, []string{s}}, true, nil
}
