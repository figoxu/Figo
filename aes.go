package Figo

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//AES/CBC/PKCS5Padding
type AesHelp struct {
	key []byte
	iv  []byte
}

func NewAesHelp(keyParam []byte, ivParam ...byte) AesHelp {
	key, iv := keyParam, keyParam
	if len(ivParam) > 0 {
		iv = ivParam
	}

	return AesHelp{
		key: key,
		iv:  iv,
	}
}

func (p *AesHelp) Encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(p.key)
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
	blockMode := cipher.NewCBCEncrypter(block, p.iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (p *AesHelp) Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, p.iv[:blockSize])
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
