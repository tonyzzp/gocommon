package gocommon

import (
	"crypto/cipher"
	"crypto/aes"
)

//AES加解密
type AES struct {
	iv    []byte
	block cipher.Block
}

func (a *AES) init(key, iv []byte) {
	a.iv = iv
	b, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	a.block = b
}

func (a *AES) check() {
	if a.block == nil {
		panic("需要先调用init进行初始化")
	}
}

func (a *AES) encrypt(data []byte) []byte {
	a.check()
	mode := cipher.NewCBCEncrypter(a.block, a.iv)
	data = padding(data, mode.BlockSize())
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return result
}

func (a *AES) decrypt(data []byte) []byte {
	a.check()
	mode := cipher.NewCBCDecrypter(a.block, a.iv)
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return unpadding(result)
}
