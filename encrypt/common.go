package encrypt

import (
	"github.com/tonyzzp/gocommon/bytesutil"
)

func padding(data []byte, blockSize int) []byte {
	l := blockSize - len(data)%blockSize
	p := bytesutil.Repeat(l, l)
	data = append(data, p...)
	return data
}

func unpadding(data []byte) []byte {
	v := data[(len(data) - 1)]
	return data[:(len(data) - int(v))]
}
