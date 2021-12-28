package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

const (
	SECRETKEY = "abcdefgh" //八字节密匙
	IV        = "joker"    //CBC模式的初始化向量
)

func testDES() {
	srcString := "timliudream"
	//加密
	cryptBytes := EncryptDES([]byte(srcString), []byte(SECRETKEY))
	fmt.Println("加密效果:")
	fmt.Println("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	fmt.Println("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println()

	//解密
	deCryptBytes := DecryptDES(cryptBytes, []byte(SECRETKEY))
	fmt.Println("解密效果:")
	fmt.Println("src :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println("dst :", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	fmt.Println()
	fmt.Println("原字串为:" + srcString)
	fmt.Println("解密字符为:" + string(deCryptBytes))
}

//使用des进行对称加密
func EncryptDES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	//2. 对最后一个明文分组进行数据填充
	src = paddingBytes(src, block.BlockSize())
	//3. 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	cbcDecrypter := cipher.NewCBCEncrypter(block, []byte(IV))
	//4. 加密连续的数据并返回
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)

	return dst
}

//使用des进行解密
func DecryptDES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(IV))
	//3. 数据块解密
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	//4. 去掉最后一组填充数据
	newBytes := unPaddingBytes(dst)
	return newBytes
}
