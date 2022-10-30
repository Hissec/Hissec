package encryptPool

import (
	"Hissec"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//密文最多比明文长16字节
type Aes struct {
}

func init() {
	regist("aes", &Aes{})
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AesEncrypt 加密函数
func aesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func aesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = pkcs7UnPadding(origData)
	return origData, nil
}
func (*Aes) Encode(data []byte) []byte {
	if len(data) < 32 {
		return data
	}
	key := data[8:24]
	iv := Hissec.BytesCombine(data[:8], data[24:32])
	ciphertext, _ := aesEncrypt(data[32:], key, iv)
	return Hissec.BytesCombine(data[:32], ciphertext)
}

func (*Aes) Decode(data []byte) ([]byte, error) {
	if len(data) < 32 {
		return data, nil
	}
	key := data[8:24]
	iv := Hissec.BytesCombine(data[:8], data[24:32])
	text, err := aesDecrypt(data[32:], key, iv)
	return Hissec.BytesCombine(data[:32], text), err
}
