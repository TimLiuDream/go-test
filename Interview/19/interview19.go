package main

import "fmt"

func GetValue(m map[int]string, id int) (value string, result bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return value, false
}
func main()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}
