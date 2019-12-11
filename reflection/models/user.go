package models

import "fmt"

//定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

//定义结构体方法
func (u User) Hello() {
	fmt.Println("Hello xuxuebiao")
}

func (u User) Hello1(name string, id int) {
	fmt.Printf("Hello %s,my name is %s and my id is %d\n", name, u.Name, id)
}
