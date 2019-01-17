package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
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
