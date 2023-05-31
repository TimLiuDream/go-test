package main

// 文件名作为参数传入
// 文件内容每行是一个 URL，文件内容会非常大，会有可能千万行
// 请爬取 url 的内容，并把结果写入到以 行号为文件名的文件中
// 单机运行，不使用中间件，尽量快并且省内存
// 直接使用将文件拆成多个小文件，每个文件 1000 行，然后并发处理
// 并发数不超过 100 个

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
)

func Read(fileName string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("open file err: ", err)
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			out <- string(line)
		}
	}()
	return out
}

func Write(fileName string) chan<- string {
	in := make(chan string)
	go func() {
		defer close(in)
		for line := range in {
			fileName := strings.Split(fileName, ".")[0] + "_" + line + ".txt"
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("create file err: ", err)
				return
			}
			defer file.Close()
			_, err = file.WriteString(line)
			if err != nil {
				fmt.Println("write file err: ", err)
				return
			}
		}
	}()
	return in
}

var fileName string

// .txt: The filename, directory name, or volume label syntax is incorrect.

func main() {
	// 文件名称作为参数，使用 flag 解析
	flag.StringVar(&fileName, "file", "", "file name")
	flag.Parse()
	ParseFile(fileName)
}

var index int

func ParseFile(fileName string) {
	// 切割文件
	// read fileName
	urls := Read(fileName)
	// 限制并发数
	limit := make(chan struct{}, 100)
	var wg sync.WaitGroup
	for url := range urls {
		limit <- struct{}{}
		wg.Add(1)
		index++
		go func(url string) {
			defer wg.Done()
			defer func() { <-limit }()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("get url err: ", err)
				return
			}
			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("read body err: ", err)
				return
			}
			// write to new file
			strco := strings.Split(fileName, ".")[0] + "_" + string(index) + ".txt"
			Write(strco) <- string(bytes)
		}(url)
	}
	wg.Wait()
}

// write some unit test for ParseFile
func TestParseFile(t *testing.T) {
	// create file
	file, err := os.Create("test.txt")
	if err != nil {
		t.Errorf("create file err: %v", err)
	}
	defer file.Close()
	// write url to file
	urls := []string{"https://www.baidu.com", "https://www.sina.com.cn", "https://www.qq.com"}
	for _, url := range urls {
		_, err := file.WriteString(url + "\n")
		if err != nil {
			t.Errorf("write url to file err: %v", err)
		}
	}
	// save file and get file name
	fileName := file.Name()
	// parse file
	ParseFile(fileName)
	// check file
	for _, url := range urls {
		fileName := strings.Split("test.txt", ".")[0] + "_" + url + ".txt"
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("open file err: %v", err)
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("read file err: %v", err)
		}
		if string(bytes) != url {
			t.Errorf("write url to file err: %v", err)
		}
	}
}
