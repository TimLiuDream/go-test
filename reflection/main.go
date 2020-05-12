package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/timliudream/golangtraining/reflection/models"
)

func main() {
	func1()
}

func func1() {
	// u := models.User{1, "bgops", 25}
	// Info(u)
	// u.Hello()
	func8()
}

func func2() {
	//注意匿名字段的初始化操作
	m := models.Manager{User: models.User{1, "biaoge", 24}, Title: "hello biao"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}

func func3() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(5256)
	fmt.Println(x)
}

//使用反射方式对结构体对象的修改有两个条件
//1.通过指针
//2.必须是可set的方法
func func4() {
	num := 123
	numv := reflect.ValueOf(&num)
	//通过Value的Elem()和SetX()方法可直接对相关的对象进行修改
	numv.Elem().SetInt(666)
	fmt.Println(num)

	u := models.User{1, "biao", 24}
	uu := reflect.ValueOf(&u)
	//Set()后面的必须是值类型
	//func (v Value) Set(x Value)
	test := models.User{2, "bgops", 2}
	testv := reflect.ValueOf(test)
	uu.Elem().Set(testv)
	fmt.Println("Change the test to u with Set(x Value)", uu)

	//此时的U已经被上面那个uu通过指针的方式修改了
	Set(&u)
	fmt.Println(u)
}

func func5() {
	u := models.User{1, "biaoge", 24}
	fmt.Println("方法调用:")
	u.Hello1("xuxuebiao", 121)

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

func func6() {
	x := 5
	y := 6
	fmt.Println(reflect.TypeOf('0'))
	fmt.Println(reflect.ValueOf('0'))
	fmt.Println(reflect.TypeOf(y - x))
	fmt.Println(reflect.ValueOf(y - x))
	fmt.Println(reflect.TypeOf(y - x + '0'))
	fmt.Println(reflect.ValueOf(y - x + '0'))
}

func func7() {
	var value = interface{}(int64(math.MaxInt64))

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Int64:
		in := int(value.(int64))
		if in > math.MaxInt32 {
			fmt.Println("234567654")
		}
	}
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

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//判断反射体值v是否是Ptr类型并且不能进行Set操作
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
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

type GTIDI interface {
	GetGTID() string
}

func func8() {
	tx := models.NewGorpTx()
	v := reflect.ValueOf(tx)
	elem := v.Elem()
	sqltxField := elem.FieldByName("sqltx")
	sqltxElem := sqltxField.Elem()
	txiField := sqltxElem.FieldByName("txi")
	txiElem := txiField.Elem()
	txiFieldType := txiField.Type()
	txiFieldKind := txiField.Kind()
	fmt.Println(txiFieldType)
	fmt.Println(txiFieldKind)

	for i := 0; i < txiElem.NumMethod(); i++ {
		fmt.Println(txiElem.Method(i).String())
	}
	txiElemI := txiElem.Interface()
	fmt.Println(txiElemI)
	if _, ok := txiField.Addr().Elem().Interface().(GTIDI); ok {
		fmt.Println("===ok ")
	}
	for i := 0; i < txiField.NumField(); i++ {
		fmt.Println(txiField.Field(i).String())
	}
	ef := txiField
	e := txiElem

	fmt.Println(ef, e)

	tx.Close()

	for i := 0; i < e.NumMethod(); i++ {
		fmt.Println(e.Method(i).String())
	}
	eValues := e.MethodByName("GetGTID").Call(nil)
	fmt.Printf("eValues: %+v\n", eValues)
}
