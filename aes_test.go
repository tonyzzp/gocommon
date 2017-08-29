package gocommon

import (
	"testing"
	"fmt"
)

func Test_aes_right(t *testing.T) {
	a := AES{}
	key := []byte{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}
	a.init(key, key)
	s := "你好helloworld!"
	result := a.encrypt([]byte(s))
	fmt.Println("加密结果", result)
	result = a.decrypt(result)
	descrypt := string(result)
	fmt.Println("解密结果", descrypt)
	if descrypt != s {
		t.FailNow()
	}
}

func Test_aes_wrong1(t *testing.T) {
	a := AES{}
	key := []byte{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4}
	a.init(key, key)
}

func Test_aes_wrong2(t *testing.T) {
	a := AES{}
	//key := []byte{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}
	a.encrypt([]byte("aa"))

}
