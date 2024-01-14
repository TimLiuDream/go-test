package main

import (
	"fmt"
	"reflect"
)

//func IsNil(x interface{}) bool {
//	if x == nil {
//		return true
//	}
//
//	return reflect.ValueOf(x).IsNil()
//}

func IsNil(x interface{}) bool {
	if x == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}

func main() {
	fmt.Printf("int IsNil: %t\n", IsNil(returnInt()))
	fmt.Printf("intPtr IsNil: %t\n", IsNil(returnIntPtr()))
	fmt.Printf("slice IsNil: %t\n", IsNil(returnSlice()))
	fmt.Printf("slicePtr IsNil: %t\n", IsNil(returnSlicePtr()))
	fmt.Printf("map IsNil: %t\n", IsNil(returnMap()))
	fmt.Printf("interface IsNil: %t\n", IsNil(returnInterface()))
	fmt.Printf("structPtr IsNil: %t\n", IsNil(returnStructPtr()))
}

func returnInt() interface{} {
	var value int
	return value
}

func returnIntPtr() interface{} {
	var value *int
	return value
}

func returnSlice() interface{} {
	var value []string
	return value
}

func returnSlicePtr() interface{} {
	var value *[]string
	return value
}

func returnMap() interface{} {
	var value map[string]struct{}
	return value
}

func returnInterface() interface{} {
	var value interface{}
	return value
}

type People struct {
	Name string
}

func returnStructPtr() interface{} {
	var value *People
	return value
}
