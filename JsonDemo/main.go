package main

import (
	"encoding/json"
	"fmt"
)

//这里的字段需要注意
//当字段是小写字母开头的话，是不可导出的，所以在序列化的时候并不能正确的序列化,即使有后面的标记也是不可以的
//字段一定要大写开头
type Result struct {
	//status int `json:"status"`
	Status int
}

func main() {
	var s string = `{"status":200}`
	r := &Result{}

	err := json.Unmarshal([]byte(s), r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
	fmt.Println(r.Status)
}
