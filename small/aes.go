// aes.go
package small

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	//"fmt"
)

var (
	ErrPaddingSize = errors.New("padding size error")
)

// 接口返回的加密数据( encryptedData )进行对称解密。 解密算法如下：
// 对称解密使用的算法为 AES-128-CBC，数据采用PKCS#7填充。
// 对称解密的目标密文为 Base64_Decode(encryptedData)。
// 对称解密秘钥 aeskey = Base64_Decode(session_key), aeskey 是16字节。
// 对称解密算法初始向量 为Base64_Decode(iv)，其中iv由数据接口返回。
// AES并没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func AesCBCDecrypt(session_key, encryptedData, iv string) (plaintext []byte,
	err error) {
	// base
	aeskey, err := base64.StdEncoding.DecodeString(session_key)
	if err != nil {
		return nil, errors.New("base64 decoding session_key err:" + err.Error())
	}
	encryptedDatabytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, errors.New("base64 decoding encryptedData err:" + err.Error())
	}
	ivbytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, errors.New("base64 decoding iv err:" + err.Error())
	}
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}
	plaintext = make([]byte, len(encryptedDatabytes))
	cipher.NewCBCDecrypter(block, ivbytes).CryptBlocks(plaintext, encryptedDatabytes)
	return PKCS5UnPadding(plaintext, block.BlockSize())
}

//https://play.golang.org/p/oZ5OwdRYLV
func PKCS5UnPadding(plaintext []byte, blockSize int) ([]byte, error) {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	//
	if unpadding >= length || unpadding > blockSize {
		return nil, ErrPaddingSize
	}
	return plaintext[:(length - unpadding)], nil
}
