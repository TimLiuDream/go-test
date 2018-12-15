package main

import (
	"fmt"
	"reflect"
)

//定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

//定义结构体方法
func (u User) Hello() {
	fmt.Println("Hello xuxuebiao")
}

func main() {
	u := User{1, "bgops", 25}
	Info(u)
	u.Hello()
}

//定义一个反射函数,参数为任意类型
func Info(o interface{}) {
	//使用反射类型获取o的Type,一个包含多个方法的interface
	t := reflect.TypeOf(o)
	//打印类型o的名称
	fmt.Println("type:", t.Name())

	//使用反射类型获取o的Value,一个空的结构体
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	//t.NumField()打印结构体o的字段个数(Id,Name,Age共三个)
	for i := 0; i < t.NumField(); i++ {
		//根据结构体的下标i来获取结构体某个字段,并返回一个新的结构体
		/**
		type StructField struct {
			Name string
			PkgPath string
			Type      Type
			Tag       StructTag
			Offset    uintptr
			Index     []int
			Anonymous bool
		}
		**/
		f := t.Field(i)

		//使用结构体方法v.Field(i)根据下标i获取字段Value(Id,Name,Age)
		//在根据Value的Interface()方法获取当前的value的值(interface类型)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	//使用t.NumMethod()获取所有结构体类型的方法个数(只有Hello()一个方法)
	//接口Type的方法NumMethod() int
	for i := 0; i < t.NumMethod(); i++ {
		//使用t.Method(i)指定方法下标获取方法对象。返回一个Method结构体
		//Method(int) Method
		m := t.Method(i)
		//打印Method结构体的相关属性
		/*
		type Method struct {
			  Name    string
			  PkgPath string
			  Type    Type
			  Func    Value
			  Index   int
		}
		*/
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}
