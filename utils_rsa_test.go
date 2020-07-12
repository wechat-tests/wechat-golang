package wxxx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	res, err := RsaEncrypt(data)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	fmt.Printf(" rsa res :len= %d >>  %s\n", len(res), ArrToHexStrWithSp(res, ""))
}
