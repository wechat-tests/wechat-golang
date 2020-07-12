package wxxx

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

//noinspection ALL
func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//noinspection ALL
func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//noinspection ALL
func AbsInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

//noinspection ALL
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

//noinspection ALL
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

//noinspection ALL
func AbsInt64(a int64) int64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

//convert uint8 array to string
//noinspection ALL
func ArrToHexStr(arr []uint8) string {
	return ArrToHexStrWithSp(arr, " ")
}

//noinspection ALL
func ArrToHexStrWithSp(arr []uint8, sp string) string {
	if arr == nil || len(arr) == 0 {
		return ""
	}
	builder := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			builder.WriteString(fmt.Sprintf("%02X", arr[i]))
		} else {
			builder.WriteString(fmt.Sprintf("%02X%s", arr[i], sp))
		}
	}
	return builder.String()
}

//big ending
//noinspection ALL
func Uint16ToArrBE(v uint16) []uint8 {
	return []uint8{uint8(v>>8) & 0xFF, uint8(v) & 0xFF}
}

//big ending
//noinspection ALL
func ArrToUint16BE(arr []uint8) uint16 {
	if len(arr) != 2 {
		return 0
	}
	return (uint16(arr[0]) << 8) | uint16(arr[1])
}

//big ending
//noinspection ALL
func Int16ToArrBE(v uint16) []uint8 {
	return []uint8{uint8(v>>8) & 0xFF, uint8(v) & 0xFF}
}

//big ending
//noinspection ALL
func ArrToInt16BE(arr []uint8) int16 {
	if len(arr) != 2 {
		return 0
	}
	return (int16(arr[0]) << 8) | int16(arr[1])
}

//big ending
//noinspection ALL
func Uint32ToArrBE(v uint32) []uint8 {
	return []uint8{uint8(v>>24) & 0xFF, uint8(v>>16) & 0xFF, uint8(v>>8) & 0xFF, uint8(v) & 0xFF}
}

//big ending
//noinspection ALL
func ArrToUint32BE(arr []uint8) uint32 {
	if len(arr) != 4 {
		return 0
	}
	return (uint32(arr[0]) << 24) | (uint32(arr[1]) << 16) | (uint32(arr[2]) << 8) | uint32(arr[3])
}

//big ending
//noinspection ALL
func Int32ToArrBE(v int32) []uint8 {
	return []uint8{uint8(v>>24) & 0xFF, uint8(v>>16) & 0xFF, uint8(v>>8) & 0xFF, uint8(v) & 0xFF}
}

//big ending
//noinspection ALL
func UInt32ToArrBE(v uint32) []uint8 {
	return []uint8{uint8(v>>24) & 0xFF, uint8(v>>16) & 0xFF, uint8(v>>8) & 0xFF, uint8(v) & 0xFF}
}

//big ending
//noinspection ALL
func ArrayToInt32BE(arr []uint8) int32 {
	if len(arr) != 4 {
		return 0
	}
	return (int32(arr[0]) << 24) | (int32(arr[1]) << 16) | (int32(arr[2]) << 8) | int32(arr[3])
}

//big ending
//noinspection ALL
func Uint64ToArrBE(v uint64) []uint8 {
	return []uint8{
		uint8(v>>(8*7)) & 0xFF,
		uint8(v>>(8*6)) & 0xFF,
		uint8(v>>(8*5)) & 0xFF,
		uint8(v>>(8*4)) & 0xFF,
		uint8(v>>(8*3)) & 0xFF,
		uint8(v>>(8*2)) & 0xFF,
		uint8(v>>(8*1)) & 0xFF,
		uint8(v>>(8*0)) & 0xFF,
	}
}

//big ending
//noinspection ALL
func ArrToUint64BE(arr []uint8) uint64 {
	if len(arr) != 8 {
		return 0
	}
	return (uint64(arr[0]) << (8 * 7)) |
		(uint64(arr[1]) << (8 * 6)) |
		(uint64(arr[2]) << (8 * 5)) |
		(uint64(arr[3]) << (8 * 4)) |
		(uint64(arr[4]) << (8 * 3)) |
		(uint64(arr[5]) << (8 * 2)) |
		(uint64(arr[6]) << (8 * 1)) |
		(uint64(arr[7]) << (8 * 0))
}

//noinspection ALL
func GetMd5Hex(data []byte) string {
	md5Bytes := md5.Sum(data)
	return hex.EncodeToString(md5Bytes[:])
}

//noinspection ALL
func Md5Hex(str string) string {
	return GetMd5Hex([]byte(str))
}

func Md5Bytes(str string) []byte {
	md5Bytes := md5.Sum([]byte(str))
	return md5Bytes[:]
}
