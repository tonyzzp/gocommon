package encrypt

import (
	"bytes"
	"fmt"
	"github.com/tonyzzp/gocommon/bytesutil"
	"math/rand"
	"testing"
)

func Test_aes(t *testing.T) {
	newData := func() []byte {
		size := rand.Intn(5000)
		data := make([]byte, size)
		for i := 0; i < size; i++ {
			data[i] = byte(rand.Intn(256))
		}
		return data
	}

	for i := 0; i < 20; i++ {
		a := AES{}
		a.Init(bytesutil.Repeat(5, 32), bytesutil.Repeat(3, 16), Mode_Encrypt)
		data := newData()
		fmt.Println("")
		fmt.Println("原始数据", data)
		r := bytes.NewReader(data)
		buf := make([]byte, 100)
		var result []byte
		for {
			l, _ := r.Read(buf)
			if l <= 0 {
				break
			} else {
				code := a.Update(buf[:l])
				result = append(result, code...)
			}
		}
		result = append(result, a.Dofinal()...)
		fmt.Println("加密数据", result)

		a.Init(bytesutil.Repeat(5, 32), bytesutil.Repeat(3, 16), Mode_Decrypt)
		r = bytes.NewReader(result)
		var result2 []byte
		for {
			l, _ := r.Read(buf)
			if l <= 0 {
				break
			} else {
				code := a.Update(buf[:l])
				result2 = append(result2, code...)
			}
		}
		result2 = append(result2, a.Dofinal()...)
		fmt.Println("解密数据", result2)

		flag := bytes.Equal(result2, data)
		fmt.Println("结果正确", flag)
		if !flag {
			t.Fail()
		}
	}
}
