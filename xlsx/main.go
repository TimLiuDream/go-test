package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

var (
	inFile = "/Users/chain/Downloads/student1.xlsx"
	// outFile = "/Users/chain/Downloads/out_student.xlsx"
	outFile = "F:\\out_student.xlsx"
)

type Student struct {
	Name   string
	age    int
	Phone  string
	Gender string
	Mail   string
}

func Export() {
	//与项目的写法不同的是
	//项目用的是指针对象
	//这里用的是普通对象
	// 而项目中会出现头乱码的情况
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1") // sheet的名字不能是中文
	if err != nil {
		fmt.Printf(err.Error())
	}
	stus := getStudents()
	headRow := sheet.AddRow()
	aCell := headRow.AddCell()
	aCell.Value = "不会有乱码吧！！"

	//add data
	for _, stu := range stus {
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = stu.Name

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(stu.age)

		phoneCell := row.AddCell()
		phoneCell.Value = stu.Phone

		genderCell := row.AddCell()
		genderCell.Value = stu.Gender

		mailCell := row.AddCell()
		mailCell.Value = stu.Mail
	}
	err = file.Save(outFile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")
}

func getStudents() []Student {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = "name" + strconv.Itoa(i+1)
		stu.Mail = stu.Name + "@chairis.cn"
		stu.Phone = "1380013800\n" + strconv.Itoa(i)
		stu.age = 20
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}

func main() {
	Export()
}
