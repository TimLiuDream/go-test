package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	path := "C:\\Users\\TimLiu\\go\\src\\github.com\\timliudream\\golangtraining\\read_file\\test.txt"
	time1 := time.Now()
	fmt.Println("------------------------")
	_ = readFile1(path)
	time2 := time.Now()
	fmt.Println(time2.Sub(time1))
	fmt.Println("------------------------")

	fmt.Println("------------------------")
	_ = readFile2(path)
	time3 := time.Now()
	fmt.Println(time3.Sub(time2))
	fmt.Println("------------------------")

	fmt.Println("------------------------")
	_ = readFile3(path)
	time4 := time.Now()
	fmt.Println(time4.Sub(time3))
	fmt.Println("------------------------")
}

// 使用ioutil的ReadFile方法
func readFile1(path string) []byte {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bs
}

// 使用os的OpenFile方法
func readFile2(path string) []byte {
	f, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs2, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bs2
}

// 使用os的Open方法
func readFile3(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs3, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bs3
}
