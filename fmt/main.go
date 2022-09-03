package main

import "fmt"

func main() {
	func1()
}

func func1(a ...int) {
	fmt.Printf("%#v\n", a)
}

func func2() {
	s := "../../usage-guide/ones-project/guan-li-gong-zuo-xiang/guan-lian-guan-xi-lei-xing"
	fmt.Printf("%s\n", s)

	i := 5
	j, err := fmt.Sscanf(" ../../usage-guide/ones-project/guan-li-gong-zuo-xiang/guan-lian-guan-xi-lei-xing ", "%5s%d", &s, &i)
	if err != nil {
		panic(err)
	}
	fmt.Println(j)
	fmt.Println(s)
}

func func3() {
	s := "{\"condition_groups\":[[{\"field_uuid\":\"field024\",\"operate\":{\"label\":\"filter.addQueryContent.include\",\"operate_id\":\"include\",\"predicate\":\"in\"},\"value\":[\"%s\",\"%s\"]}]]}"
	placeholder := []interface{}{"11111", "222222"}
	fmt.Println(fmt.Sprintf(s, placeholder...))
	// s := "{\"condition_groups\":[[{\"field_uuid\":\"field024\",\"operate\":{\"label\":\"filter.addQueryContent.include\",\"operate_id\":\"include\",\"predicate\":\"in\"},\"value\":[\"%s\",\"%s\"]}]]}"
	// placeholder := []interface{}{"11111", "222222"}
	// fmt.Println(fmt.Sprintf(s, placeholder...))
}

type integer int

func (i integer) String() string {
	return "hello"
}

func func4() {
	fmt.Println(integer(1))
}

// 0 开头，表明是八进制，但八进制最大的数字是 7，因此编译不通过
func func5() {
	fmt.Println(0
	9)
}
