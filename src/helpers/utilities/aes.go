package utilities

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Aes struct {
	Key []byte
}

func (self Aes) AesEncrypt(input string) (string, error) {
	origData := []byte(input)
	block, err := aes.NewCipher(self.Key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = self.pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, self.Key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (self Aes) AesDecrypt(input string) ([]byte, error) {
	crypted, arserr := base64.StdEncoding.DecodeString(input)
	if arserr != nil {
		return nil, arserr
	}

	block, err := aes.NewCipher(self.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, self.Key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = self.pkcs5UnPadding(origData)
	return origData, nil
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
