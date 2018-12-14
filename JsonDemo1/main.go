package main

import (
	"encoding/json"
	"fmt"
)

type Seq struct {
	Num  int
	Name string
}
type Info struct {
	Id    int
	Value string
}

func main() {
	nA := Seq{1, "11"}
	nB := Info{101, "aaa"}
	//方法1
	anyMarshal(struct {
		*Seq
		*Info
	}{&nA, &nB})
	//方法2
	anyMarshalNew(nA, nB)
}
func anyMarshal(msg interface{}) {
	mData, _ := json.Marshal(msg)
	fmt.Println(mData)
	fmt.Println(string(mData))
}
func anyMarshalNew(nA Seq, nB interface{}) {
	nF := struct {
		*Seq
		Any interface{}
	}{&nA, &nB}
	mData, _ := json.Marshal(nF)
	fmt.Println(mData)
	fmt.Println(string(mData))
}
