package main

import (
	"./encrypt"
	"bytes"
	"fmt"
)

func main() {
	aes := encrypt.AES{}
	aes.Init([]byte{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}, bytes.Repeat([]byte{byte(1)}, 16))
	s := "asdf"
	result := aes.Encrypt([]byte(s))
	result = aes.Decrypt(result)
	fmt.Println(string(result))
}
