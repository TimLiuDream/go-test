package main

import "bytes"

//填充明文最后一个分组工具方法
//src - 原始数据
//blockSize - 每个分组的数据长度
func paddingBytes(src []byte, blockSize int) []byte {
	//1.求出最后一个分组要填充多个字节
	padding := blockSize - len(src)%blockSize
	//2.创建新的切片，切片的字节数为填充的字节数，并初始化化，每个字节的值为填充的字节数
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	//3.将创建出的新切片和原始数据进行连接
	newBytes := append(src, padBytes...)

	//4.返回新的字符串
	return newBytes
}

//删除密文末尾分组填充的工具方法
func unPaddingBytes(src []byte) []byte {
	//1.求出要处理的切片的长度
	l := len(src)
	//2.取出最后一个字符，得到其整型值
	n := int(src[l-1])

	//3.将切片末尾的number个字节删除
	return src[:l-n]
}
