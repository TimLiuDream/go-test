package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	func1()
	// fileName := "test.txt"
	// dstFile, err := os.Create(fileName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// defer dstFile.Close()
	// s := "hello \n world!"
	// dstFile.WriteString(s + "\n")
	// path, err := filepath.Abs(dstFile.Name())
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(path)
}

func func1() {
	file, _ := os.Open("/Users/tim/Downloads/WX20200112-220316@2x.png")
	targetFilePath := "/Users/tim/Downloads/WX20200112-220316@2x_1.png"
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(targetFilePath, bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// fInfo, _ := file.Stat()
	// size := fInfo.Size()
	// buf := make([]byte, size)
	// fReader := bufio.NewReader(file)
	// fReader.Read(buf)

	// bytes, err := ioutil.ReadAll(fReader)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = ioutil.WriteFile(targetFilePath, bytes, 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// info2, _ := file.Stat()
	// fmt.Println(info2.Name())

	// abs, err := filepath.Abs(filepath.Dir(info2.Name()))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(abs)
}
