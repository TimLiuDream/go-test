package main

import (
	"fmt"
)

type Cat struct {
	name           string
	scientificName string
	category       string
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
	fmt.Printf("the cat:%s\n", cat)
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category:%s,name:%q)", cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster")
	fmt.Printf("the cat:%s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("the cat:%s\n", cat)

	//cat.name="little pig"
	//fmt.Printf("the cat:%s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName()string
	}

	cat1:=interface{}(cat)
	cat2:=interface{}(&cat)
	fmt.Print(cat1)
	fmt.Print(cat2)
	_,ok:=interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet :%v\n",ok)
	_,ok=interface{}(&cat).(Pet)
	fmt.Printf("Cat implements interface Pet :%v\n",ok)
}
