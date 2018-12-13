package main

import (
	"sync/atomic"
	"fmt"
	"errors"
	"io"
	"os"
	"reflect"
)

func main() {
	//1
	var box atomic.Value
	fmt.Println("copy box to box2")
	box2 := box //原子值在真正使用前可以被复制
	v1 := [...]int{1, 2, 3}
	fmt.Printf("store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("the value load from box is %v\n", box.Load())
	fmt.Printf("the value load from box2 is %v\n", box2.Load())
	fmt.Println()

	//2
	v2 := "123"
	fmt.Printf("store %q to box2.\n", v2)
	box2.Store(v2) //这里不会引发panic
	fmt.Printf("the value load from box is %v\n", box.Load())
	fmt.Printf("the value load from box2 is %q\n", box2.Load())
	fmt.Println()

	//3
	fmt.Println("store box to box3.")
	box3 := box //原子值在真正使用后不应该被复制
	fmt.Printf("the value load from box3 is %v\n", box3.Load())
	v3 := 123
	fmt.Printf("store %d to box3\n", v3)
	box3.Store(v3)
	fmt.Printf("the value load from box3 is %v\n", box3.Load())
	_ = box3
	fmt.Println()

	//4
	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("store an error with message %q to box4\n", v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("store a value of the same type to box4")
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Printf("store a value of type %T that implements error interface to box4.\n", v42)
		//box4.Store(v42) // 这里会引发一个panic，报告存储值的类型不一致。
	}
	fmt.Println()

	//5
	box5, err := NewAtomicValue(v4)
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}
	fmt.Printf("the legal type in box5 is %s\n", box5.TypeOfValue())

	// 示例6。
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)
	v6[1] = 4 // 注意，此处的操作不是并发安全的！
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	// 正确的做法如下。
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box6.Store(replica)
	}
	fmt.Printf("Store %v to box6.\n", v6)
	store(v6)
	v6[2] = 5 // 此处的操作是安全的。
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
}

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}

func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}
