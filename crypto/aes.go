package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

const (
	AESSECRETKEY = "abcdefgh12345678" //十六字节密匙
	AESIV        = "aaaabbbb12345678" //CBC模式的初始化向量：与key等长：十六字节
)

func TestAES() {
	fmt.Println("测试AES对称加密:")
	srcString := "timliudream"
	//加密
	cryptBytes := EncryptAES([]byte(srcString), []byte(AESSECRETKEY))
	fmt.Println("加密效果:")
	fmt.Println("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	fmt.Println("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println()

	//解密
	deCryptBytes := DecryptAES(cryptBytes, []byte(AESSECRETKEY))
	fmt.Println("解密效果:")
	fmt.Println("src bytes:", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println("dst bytes:", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	fmt.Println()
	fmt.Println("原字串为:" + srcString)
	fmt.Println("解密字符为:" + string(deCryptBytes))
}

//使用aes进行对称加密
func EncryptAES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	//2. 对最后一个明文分组进行数据填充
	src = paddingBytes(src, block.BlockSize())
	//3. 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	cbcDecrypter := cipher.NewCBCEncrypter(block, []byte(AESIV))
	//4. 加密连续的数据并返回
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)

	return dst
}

//使用aes进行解密
func DecryptAES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(AESIV))
	//3. 数据块解密
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	//4. 去掉最后一组填充数据
	newBytes := unPaddingBytes(dst)
	return newBytes
}
