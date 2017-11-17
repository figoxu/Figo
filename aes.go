package Figo

import (
	"crypto/aes"
	"bytes"
	"crypto/cipher"
)

type AesHelp struct {
	pwd []byte
}

func NewAesHelp(pwd []byte) AesHelp {
	return AesHelp{
		pwd: pwd,
	}
}

func (p *AesHelp) encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(p.pwd)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	PKCS5Padding := func(ciphertext []byte, blockSize int) []byte {
		padding := blockSize - len(ciphertext)%blockSize
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(ciphertext, padtext...)
	}
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, p.pwd[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (p *AesHelp) decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(p.pwd)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, p.pwd[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	PKCS5UnPadding := func(origData []byte) []byte {
		length := len(origData)
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)]
	}
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
