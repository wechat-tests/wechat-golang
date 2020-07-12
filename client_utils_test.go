package wxxx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_RQTSign(t *testing.T) {
	cli := newClient(&defaultHttpCli{})
	sign, err := cli.RQTSign([]byte{0x01, 0x02, 0x03})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	fmt.Printf("sign :%X\n", sign)
	if !assert.Equal(t, sign, uint32(0x21191A11)) {
		t.FailNow()
	}
	sign, err = cli.RQTSign([]byte{0x01, 0x02, 0x03, 0x04})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	fmt.Printf("sign :%X\n", sign)
	if !assert.Equal(t, sign, uint32(0x213D3B33)) {
		t.FailNow()
	}
}
