package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

const (
	TDESSECRETKEY = "abcdefgh12345678ABCDEFGH" //三组八字节密匙，即24字节
	TDESIV        = "12345678"                 //CBC模式的初始化向量，8字节
)

func testTDES() {
	fmt.Println("测试TDES对称加密:")
	srcString := "timliudream"
	//加密
	cryptBytes := EncryptTDES([]byte(srcString), []byte(TDESSECRETKEY))
	fmt.Println("加密效果:")
	fmt.Println("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	fmt.Println("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println()

	//解密
	deCryptBytes := DecryptTDES(cryptBytes, []byte(TDESSECRETKEY))
	fmt.Println("解密效果:")
	fmt.Println("src bytes:", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	fmt.Println("dst bytes:", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	fmt.Println()
	fmt.Println("原字串为:" + srcString)
	fmt.Println("解密字符为:" + string(deCryptBytes))
}

// 使用3des进行对称加密
func EncryptTDES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	//2. 对最后一个明文分组进行数据填充，和DES的算法一样
	src = paddingBytes(src, block.BlockSize())
	//3. 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	cbcDecrypter := cipher.NewCBCEncrypter(block, []byte(TDESIV))
	//4. 加密连续的数据并返回
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)

	return dst
}

// 使用3des进行解密
func DecryptTDES(src, key []byte) []byte {
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(TDESIV))
	//3. 数据块解密
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	//4. 去掉最后一组填充数据，和DES的删除算法一样
	newBytes := unPaddingBytes(dst)
	return newBytes
}
