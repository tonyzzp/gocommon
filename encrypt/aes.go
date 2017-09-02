package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

type CipherMode int

const (
	Mode_Encrypt CipherMode = iota
	Mode_Decrypt
)

// aes加解密
type AES struct {
	block      cipher.Block
	data       []byte
	mode       cipher.BlockMode
	cipherMode CipherMode
}

//初始化密钥和加解密模式
func (a *AES) Init(key, iv []byte, mode CipherMode) {
	b, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	a.cipherMode = mode
	a.block = b
	if mode == Mode_Encrypt {
		a.mode = cipher.NewCBCEncrypter(a.block, iv)
	} else if mode == Mode_Decrypt {
		a.mode = cipher.NewCBCDecrypter(a.block, iv)
	} else {
		panic("mode取值只能是 Mode_Encrypt/Mode_Decrypt")
	}
	a.data = nil
}

func (a *AES) check() {
	if a.block == nil {
		panic("需要先调用init进行初始化")
	}
}

// 传入需要加解密的数据，返回块加解密结果。如果不够一块的大小，则会回返nil
func (a *AES) Update(data []byte) []byte {
	a.check()
	bs := a.block.BlockSize()
	a.data = append(a.data, data...)
	l := len(a.data)
	if l < bs*2 {
		return nil
	} else {
		remain := l%bs + bs
		count := l - remain
		result := make([]byte, count)
		a.mode.CryptBlocks(result, a.data[:count])
		a.data = a.data[count:]
		return result
	}
}

// 返回最后一块加解密的结果。如果使用过Update方法，则此方法必须要调用
func (a *AES) Dofinal() []byte {
	a.check()
	bs := a.block.BlockSize()
	data := a.data
	if a.cipherMode == Mode_Encrypt {
		data = padding(a.data, bs)
	}
	result := make([]byte, len(data))
	a.mode.CryptBlocks(result, data)
	a.data = nil
	if a.cipherMode == Mode_Decrypt {
		result = unpadding(result)
	}
	return result
}

// 一次性传入所有数据，进行解密或加密。
// 会清空Update传入的所有剩余数据
func (a *AES) DoAll(data []byte) []byte {
	a.check()
	bs := a.block.BlockSize()
	if a.cipherMode == Mode_Encrypt {
		data = padding(data, bs)
	}
	result := make([]byte, len(data))
	a.mode.CryptBlocks(result, data)
	if a.cipherMode == Mode_Decrypt {
		result = unpadding(result)
	}
	a.data = nil
	return result
}
