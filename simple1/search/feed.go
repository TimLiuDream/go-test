package search

import (
	"os"
	"encoding/json"
	"path"
	"runtime"
	"fmt"
)

const dataFile = "../data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//此函数读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), dataFile)
	fmt.Println(datapath)
	//打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	//当函数return时，关闭文件
	defer file.Close()

	//将文件解码到一个切片里，这个切片的每一项是一个指向Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}

func main() {
}
