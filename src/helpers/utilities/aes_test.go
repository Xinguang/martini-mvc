package utilities

import (
	"testing"
)

func Test_Aes(t *testing.T) {
	a := Aes{}
	result := a.AesEncrypt("martini-mvc")
	t.Log(result)
	origData := a.AesDecrypt(result)
	t.Log(origData)
}
