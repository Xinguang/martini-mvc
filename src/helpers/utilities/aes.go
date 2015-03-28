package utilities

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/starboychina/martini-mvc/src/config"
)

type Aes struct {
	Key []byte
}

func (self Aes) AesEncrypt(input string) string {
	origData := []byte(input)
	self.Key = []byte(config.SecretAes)
	block, err := aes.NewCipher(self.Key)
	if err != nil {
		panic(err) //系统配置错误aes -key 错误
	}
	blockSize := block.BlockSize()
	origData = self.pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, self.Key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted)
}

func (self Aes) AesDecrypt(input string) string {
	self.Key = []byte(config.SecretAes)
	crypted, arserr := base64.StdEncoding.DecodeString(input)
	if arserr != nil {
		panic(arserr)
	}

	block, err := aes.NewCipher(self.Key)
	if err != nil {
		panic(err) //系统配置错误aes -key 错误
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, self.Key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = self.pkcs5UnPadding(origData)
	return string(origData)
}
func (self Aes) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (self Aes) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
