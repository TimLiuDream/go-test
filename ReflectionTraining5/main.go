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

func (u User) Hello(name string, id int) {
	fmt.Printf("Hello %s,my name is %s and my id is %d\n", name, u.Name, id)
}

func main() {
	u := User{1, "biaoge", 24}
	fmt.Println("方法调用:")
	u.Hello("xuxuebiao", 121)

	//获取结构体类型u的Value
	v := reflect.ValueOf(u)
	//根据方法名称获取Value中的方法对象
	mv := v.MethodByName("Hello")

	//构造一个[]Value类型的变量，使用Value的Call(in []Value)方法进行动态调用method
	//这里其实相当于有一个Value类型的Slice，仅一个字段
	args := []reflect.Value{reflect.ValueOf("xuxuebiao"), reflect.ValueOf(5256)}
	fmt.Println("通过反射动态调用方法:")
	//使用Value的Call(in []Value)方法进行方法的动态调用
	//func (v Value) Call(in []Value) []Value
	//需要注意的是当v的类型不是Func的化，将会panic;同时每个输入的参数args都必须对应到Hello()方法中的每一个形参上
	mv.Call(args)

}
