package utilities

import (
	"fmt"
)

func Test_Aes() {
	a := Aes{}

	result, err := a.AesEncrypt("martini-mvc")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	origData, err := a.AesDecrypt(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(origData)
}
