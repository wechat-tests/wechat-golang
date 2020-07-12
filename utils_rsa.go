package wxxx

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

var kRsaPubKeyModules = "D153E8A2B314D2110250A0A550DDACDCD77F5801F3D1CC21CB1B477E4F2DE8697D40F10265D066BE8200876BB7135EDC74CDBC7C4428064E0CDCBE1B6B92D93CEAD69EC27126DEBDE564AAE1519ACA836AA70487346C85931273E3AA9D24A721D0B854A7FCB9DED49EE03A44C189124FBEB8B17BB1DBE47A534637777D33EEC88802CD56D0C7683A796027474FEBF237FA5BF85C044ADC63885A70388CD3696D1F2E466EB6666EC8EFE1F91BC9353F8F0EAC67CC7B3281F819A17501E15D03291A2A189F6A35592130DE2FE5ED8E3ED59F65C488391E2D9557748D4065D00CBEA74EB8CA19867C65B3E57237BAA8BF0C0F79EBFC72E78AC29621C8AD61A2B79B"

const kRsaPubKeyE = "010001"

// RsaEncrypt
func RsaEncrypt(src []byte) ([]byte, error) {
	n := big.NewInt(0)
	n.SetString(kRsaPubKeyModules, 16)
	pubKey := &rsa.PublicKey{N: n, E: 0x010001}
	var data = packageData(src, pubKey.N.BitLen()/8-11)
	cipher := make([]byte, 0, 0)
	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pubKey, d)
		if e != nil {
			return nil, e
		}
		cipher = append(cipher, c...)
	}
	return cipher, nil
}

func packageData(rawData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(rawData))
	copy(src, rawData)

	r = make([][]byte, 0)
	if len(src) <= packageSize {
		return append(r, src)
	}
	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}
	return r
}
