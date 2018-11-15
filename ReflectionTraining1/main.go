package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	//注意匿名字段的初始化操作
	m := Manager{User: User{1, "biaoge", 24}, title: "hello biao"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

}
