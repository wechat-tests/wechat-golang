package wxxx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAesEncryptECB(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	key := []byte{0x32, 0x7b, 0xd9, 0xa0, 0xed, 0x80, 0x8a, 0xc4, 0x02, 0x80, 0xcd, 0xa8, 0xf7, 0xb5, 0x26, 0x45}
	encrypted := AESCbcEncrypt(data, key)
	fmt.Printf("encrypted : len = %d , >>  %s \n", len(encrypted), ArrToHexStrWithSp(encrypted, ""))
	if !assert.Equal(t, "5b7c420244edc4c445f4f1f913020c09", strings.ToLower(ArrToHexStrWithSp(encrypted, ""))) {
		t.FailNow()
	}
	decrypted := AESCbcDecrypt(encrypted, key)
	fmt.Printf("decrypted : len = %d >> %s \n", len(decrypted), ArrToHexStrWithSp(decrypted, ""))
	if !assert.Equal(t, "010203", strings.ToLower(ArrToHexStrWithSp(decrypted, ""))) {
		t.FailNow()
	}
}
