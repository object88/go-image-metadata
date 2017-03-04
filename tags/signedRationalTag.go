package tags

import (
	"bytes"
	"strconv"

	"github.com/object88/go-image-metadata/common"
)

// SignedRational is a fractional representation of an signed number
type SignedRational struct {
	Numerator   int32
	Denominator int32
}

// SignedRationalTag holds an array of unsigned rationals.
type SignedRationalTag struct {
	BaseTag
	value []SignedRational
}

func (m *SignedRationalTag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.GetName())
	buffer.WriteString(" (")
	buffer.WriteString(common.Srational.String())
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

func readSignedRational(reader TagReader, name string, raw *RawTagData) (Tag, bool, error) {
	r := reader.GetReader()
	cur := r.GetCurrentOffset()
	r.SeekTo(int64(raw.Data))
	v := make([]SignedRational, raw.Count)
	for i := uint32(0); i < raw.Count; i++ {
		n, _ := r.ReadUint32()
		d, _ := r.ReadUint32()
		v[i] = SignedRational{Numerator: int32(n), Denominator: int32(d)}
	}
	r.SeekTo(cur)
	return &SignedRationalTag{BaseTag{name, raw.Tag, raw.Format}, v}, true, nil
}
