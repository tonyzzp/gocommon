package encrypt

import (
	"crypto/des"
	"crypto/cipher"
)

//DES加解密
type DES struct {
	iv    [8]byte
	block cipher.Block
}

func (a *DES) check() {
	if a.block == nil {
		panic("需要先调用init进行初始化")
	}
}

func (d *DES) Init(key, iv [8]byte) {
	d.iv = iv
	block, e := des.NewCipher(key[:])
	if e != nil {
		panic(e)
	}
	d.block = block
}

func (d *DES) Encrypt(data []byte) []byte {
	d.check()
	mode := cipher.NewCBCEncrypter(d.block, d.iv[:])
	data = padding(data, mode.BlockSize())
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return result
}

func (d *DES) Decrypt(data []byte) []byte {
	d.check()
	mode := cipher.NewCBCDecrypter(d.block, d.iv[:])
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return unpadding(result)
}
