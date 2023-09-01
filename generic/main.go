package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	//fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
	//fmt.Println(Contains([]int{1, 2, 3}, 3))
	DeserializeUsage()
}

type constraint interface {
	~float64 | int
}

func sumAll[T constraint](arr []T) T {
	var s T
	for _, ele := range arr {
		s += ele
	}
	return s
}

func runSumAll() {
	fmt.Println("sum: ", sumAll([]int{1, 2, 3, 5, 6}))
	fmt.Println("sum: ", sumAll([]float64{1.2, 2.1, 3.8, 5.4}))
}

func func1() {
	intSlice := []int{1, 2, 3, 4, 5}
	firstInt := First[int](intSlice) // returns 1

	println(firstInt)

	stringSlice := []string{"apple", "banana", "cherry"}
	firstString := First[string](stringSlice) // returns "apple"

	println(firstString)
}

func First[T any](items []T) T {
	return items[0]
}

func SumGenerics[T int | int16 | int32 | int64 | int8 | float32 | float64](a, b T) T {
	return a + b
}

func func2() {
	sumInt := SumGenerics[int](2, 3)

	sumFloat := SumGenerics[float32](2.5, 3.5)

	sumInt64 := SumGenerics[int64](10, 20)

	fmt.Println(sumInt)   // returns 5
	fmt.Println(sumFloat) // returns 6.0
	fmt.Println(sumInt64) // returns 30
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func Serialize[T any](data T) ([]byte, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Deserialize[T any](b []byte) (T, error) {
	buffer := bytes.Buffer{}
	buffer.Write(b)
	decoder := gob.NewDecoder(&buffer)
	var data T
	err := decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func DeserializeUsage() {
	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St.",
	}

	serialized, err := Serialize(person)
	if err != nil {
		panic(err)
	}

	deserialized, err := Deserialize[Person](serialized)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Age: %d, Address: %s", deserialized.Name, deserialized.Age, deserialized.Address)
}
