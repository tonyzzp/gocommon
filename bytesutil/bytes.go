package bytesutil

import (
	"bytes"
)

func Repeat(value int, times int) []byte {
	return bytes.Repeat([]byte{byte(value)}, times)
}
