package utilities

import (
	"testing"
)

func Test_Aes() {
	a := Aes{[]byte("yVHlew1jDlZpJ/zSbJ8JPjIc2dBeoLny")}

	result, err := a.AesEncrypt("martini-mvc")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	origData, err := a.AesDecrypt(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
