package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// UnsignedRational is a fractional representation of an unsigned number
type UnsignedRational struct {
	Numerator   uint32
	Denominator uint32
}

// UnsignedRationalTag holds an array of unsigned rationals.
type UnsignedRationalTag struct {
	BaseTag
	value []UnsignedRational
}

func (m *UnsignedRationalTag) String() string {
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

func readUnsignedRational(reader TagReader, name string, raw *RawTagData) (Tag, bool, error) {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(raw.Data))
	v := make([]UnsignedRational, raw.Count)
	for i := uint32(0); i < raw.Count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = UnsignedRational{Numerator: n, Denominator: d}
	}
	r.SeekTo(cur)
	return &UnsignedRationalTag{BaseTag{name, raw.Tag, raw.Format}, v}, true, nil
}
