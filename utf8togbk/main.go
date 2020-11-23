package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	ss := "銆愮敤鎴锋晠浜嬫弿杩般€戯細浣滀负___锛屾垜甯屾湜___锛屼互渚縚__銆傘€愰獙鏀舵爣鍑嗏憼銆戯細given__中__锛寃hen____锛宼hen____銆傞獙鏀舵柟娉曪細"
	gbk, err := Utf8ToGbk([]byte(ss))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(gbk))
	}
}
