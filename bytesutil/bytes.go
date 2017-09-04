package bytesutil

import (
	"bytes"
)

func Repeat(value int, times int) []byte {
	return bytes.Repeat([]byte{byte(value)}, times)
}

func BytesToInt32(b []byte) int32 {
	i := int32(int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3]))
	return i
}

func Int32ToBytes(i int32) []byte {
	b := []byte{
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
	return b
}
