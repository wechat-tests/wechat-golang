package wxxx

import (
	"bytes"
	"math"
)

// encode uint32 value to variant
func VariantEncode(v uint32) []byte {
	out := bytes.Buffer{}
	tmp := v
	for tmp >= 0x80 {
		out.WriteByte((byte)(0x80 + tmp&0x7f))
		tmp = tmp >> 7
	}
	out.WriteByte(byte(tmp))
	return out.Bytes()
}

// decode uint32 value from variant byte[5]
func VariantDecode(buf []byte, offset int) (n int, v uint32) {
	v = uint32(0)
	min := int(math.Min(float64(len(buf)-offset), 5))
	var i int
	for i = 0; i < min; i++ {
		v |= uint32(buf[offset+i]&0x7f) << (7 * i)
		if buf[offset+i]&0x80 == 0 {
			break
		}
	}
	n = i + 1
	return n, v
}
