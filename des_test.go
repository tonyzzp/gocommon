package gocommon

import (
	"testing"
	"fmt"
)

func Test_des_right(t *testing.T) {
	d := DES{}
	d.init([8]byte{1, 2, 3, 4, 5, 6, 7, 8}, [8]byte{0, 1, 2, 3, 4, 5, 6, 7})
	src := "hello des.中文△＠→●★※＿◆"
	result := d.encrypt([]byte(src))
	fmt.Println("加密结果", result)
	decrypt := string(d.decrypt(result))
	fmt.Println("解密结果", decrypt)
	if decrypt != src {
		t.Fail()
	}
}
