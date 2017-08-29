package gocommon

import "bytes"

func padding(data []byte, blockSize int) []byte {
	l := blockSize - len(data)%blockSize
	p := bytes.Repeat([]byte{byte(l)}, l)
	data = append(data, p...)
	return data
}

func unpadding(data []byte) []byte {
	v := data[(len(data) - 1)]
	return data[:(len(data) - int(v))]
}
