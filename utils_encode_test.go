package wxxx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariantEncode(t *testing.T) {
	v := []uint32{25, 502, 174, 0x01004567}
	vAfterEncode := []string{"19", "F6 03", "AE 01", "E7 8A 81 08"}
	for i := range v {
		encodeV := VariantEncode(v[i])
		hexStr := ArrToHexStr(encodeV)
		fmt.Printf(" number %d variant encode : %s\n", v[i], hexStr)
		if !assert.Equal(t, vAfterEncode[i], hexStr) {
			t.FailNow()
		}
	}
}
func TestVariantDecode(t *testing.T) {
	encodeV := [][]byte{{0x19, 0x00, 0x00, 0x00}, {0xF6, 0x03}, {0xE7, 0x8A, 0x81, 0x08}}
	rawV := []uint32{25, 502, 0x01004567}

	for i := range encodeV {
		n, v := VariantDecode(encodeV[i], 0)
		hexStr := ArrToHexStr(encodeV[i][:])
		fmt.Printf(" variant encode %s after decode is : %d , len = %d \n", hexStr, v, n)
		if !assert.Equal(t, rawV[i], v) {
			t.FailNow()
		}
	}
}
