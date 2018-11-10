package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Test struct {
	Id     int
	Year   int
	Amount int
	Num    int
}

func (Test) TableName() string {
	return "test"
}

type UserTest struct {
	gorm.Model
	Name string
}

func (UserTest) TableName() string {
	return "user_test"
}

func main() {
	db, err := gorm.Open("mysql", "root:lft8306488@tcp(localhost:3306)/todo_list_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("orm is disconnect")
		return
	}

	db.SingularTable(true) //全局设置表名不可以为复数形式
	//判断表是否存在，不在就创建，在的话就不管
	if !db.HasTable(&Test{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Test{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&UserTest{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserTest{}).Error; err != nil {
			panic(err)
		}
	}

	var test Test
	db.First(&test)
	fmt.Println(test.Year)

	userTest := UserTest{Name: "TimLiu"}
	if err := db.Create(&userTest).Error; err != nil {
		panic(err)
	}

	if err := db.Where(&UserTest{Name: "TimLiu"}).Delete(UserTest{}).Error; err != nil {
		panic(err)
	}

	var utest UserTest
	db.First(&utest)
	fmt.Println(utest.Name)

	db.Model(&test).Where(&Test{Id:1}).Update(Test{Year: 1991})
}
