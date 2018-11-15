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

//使用反射方式对结构体对象的修改有两个条件
//1.通过指针
//2.必须是可set的方法
func main() {
	num := 123
	numv := reflect.ValueOf(&num)
	//通过Value的Elem()和SetX()方法可直接对相关的对象进行修改
	numv.Elem().SetInt(666)
	fmt.Println(num)

	u := User{1, "biao", 24}
	uu := reflect.ValueOf(&u)
	//Set()后面的必须是值类型
	//func (v Value) Set(x Value)
	test := User{2, "bgops", 2}
	testv := reflect.ValueOf(test)
	uu.Elem().Set(testv)
	fmt.Println("Change the test to u with Set(x Value)", uu)

	//此时的U已经被上面那个uu通过指针的方式修改了
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//判断反射体值v是否是Ptr类型并且不能进行Set操作
	if v.Kind() == reflect.Ptr && ! v.Elem().CanSet() {
		fmt.Println("xxx")
		return
		//初始化对象修改后的返回值(可接受v或v的指针)
	} else {
		v = v.Elem()
	}
	//按照结构体对象的名称进行查找filed，并判断类型是否为string，然后进行Set
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYBY")
	}
}
