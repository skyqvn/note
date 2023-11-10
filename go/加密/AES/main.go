//package main
//
//import (
//	"crypto/aes"
//	"encoding/hex"
//	"fmt"
//)
//
////func main() {
////	// cipher key
////	key := "thisis32bitlongpassphraseimusing"
////	// plaintext
////	pt := "This is a secret"
////	c := EncryptAES([]byte(key), pt)
////	// plaintext
////	fmt.Println(pt)
////	// ciphertext
////	fmt.Println(c)
////	// decrypt
////	DecryptAES([]byte(key), c)
////}
//
//func main() {
//	//PKCS7Padding，假设数据长度需要填充n(n>0)个字节才对齐，那么填充n个字节，每个字节都是n;如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小。
//
//	// cipher key
//	key := "thisis32bitlongpassphraseimusing"
//	// plaintext
//	pt := "This is a secret"
//	pu := "Another secret.."
//	c := EncryptAES([]byte(key), pt)
//	d := EncryptAES([]byte(key), pu)
//	e := EncryptAES([]byte(key), pt+pu)
//	// ciphertext
//	fmt.Println(c, d, e)
//	// decrypt
//	DecryptAES([]byte(key), c)
//	DecryptAES([]byte(key), d)
//	DecryptAES([]byte(key), e)
//}
//
//func EncryptAES(key []byte, plaintext string) string {
//	c, err := aes.NewCipher(key)
//	CheckError(err)
//	out := make([]byte, len(plaintext))
//	c.Encrypt(out, []byte(plaintext))
//	return hex.EncodeToString(out)
//}
//
//func DecryptAES(key []byte, ct string) {
//	ciphertext, _ := hex.DecodeString(ct)
//	c, err := aes.NewCipher(key)
//	CheckError(err)
//	pt := make([]byte, len(ciphertext))
//	c.Decrypt(pt, ciphertext)
//	s := string(pt[:])
//	fmt.Println("DECRYPTED:", s)
//}
//
//func CheckError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AEC/CBC/PKCS7Padding

func main() {
	key := "12345678901234567890123456789012" // 长度为 16、 24 或 32
	iv := "1234567890123456" // 长度固定为 aes.BlockSize ，16位
	
	s := "任何一种技艺达到完美~都会令人无法抗拒~"
	
	result, err := Encrypt(s, key, iv)
	if err != nil {
		fmt.Print(err)
	}
	
	fmt.Println(result)
	
	raw, err := Decrypt(result, key, iv)
	if err != nil {
		fmt.Print(err)
	}
	
	fmt.Println(raw)
}

// Encrypt 加密
//
// plainText: 加密目标字符串
// key: 加密Key
// iv: 加密iv(AES时固定为16位)
func Encrypt(plainText string, key string, iv string) (string, error) {
	data, err := aesCBCEncrypt([]byte(plainText), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	
	return base64.StdEncoding.EncodeToString(data), nil
}

// Decrypt 解密
//
// cipherText: 解密目标字符串
// key: 加密Key
// iv: 加密iv(AES时固定为16位)
func Decrypt(cipherText string, key string, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	
	dnData, err := aesCBCDecrypt(data, []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	
	return string(dnData), nil
}

// aesCBCEncrypt AES/CBC/PKCS7Padding 加密
func aesCBCEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	
	// PKCS7 填充
	plaintext = paddingPKCS7(plaintext, aes.BlockSize)
	
	// CBC 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(plaintext, plaintext)
	
	return plaintext, nil
}

// aesCBCDecrypt AES/CBC/PKCS7Padding 解密
func aesCBCDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	
	// CBC 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)//可多次调用
	
	// PKCS7 反填充
	result := unPaddingPKCS7(ciphertext)
	return result, nil
}

// PKCS7 填充
func paddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(plaintext, paddingText...)
}

// PKCS7 反填充
func unPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	unPadding := int(s[length-1])
	return s[:(length - unPadding)]
}
