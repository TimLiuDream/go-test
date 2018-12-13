package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	var dog1 *Dog
	fmt.Println("the first dog is nil")
	dog2 := dog1
	fmt.Println("the second dog is nil")
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("the pet is nil")
	} else {
		fmt.Println("the pet is not nil")
	}
	fmt.Printf("the type of pet is %T.\n", pet)
	fmt.Printf("the type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("the type of second dog is %T.\n", dog2)
	fmt.Println()

	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("the pet is nil")
	} else {
		fmt.Println("the pet is not nil")
	}
}
