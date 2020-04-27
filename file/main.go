package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	getFileMimetype()
}

func func1() {
	file, _ := os.Open("/Users/tim/Downloads/test")
	fmt.Println("open file")
	time.Sleep(5 * time.Second)
	defer file.Close()
	targetFilePath := "/Users/tim/Downloads/test1"
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(targetFilePath, bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func func2() {
	filePath := "testdata/test.zip"
	file, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("open file")
	defer file.Close()
	fmt.Println("file exist!")
}

// 获取文件的mimetype
func getFileMimetype() {
	filePath := "testdata/test.html"
	mime, err := mimetype.DetectFile(filePath)
	fmt.Printf("mimetype: %s ,\nextension: %s ,\nerr: %+v\n", mime.String(), mime.Extension(), err)
}
