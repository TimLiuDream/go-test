package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct{
	gorm.Model
	Code string
	Price uint
}

func main()  {
	//连接数据库，对于sqlite没有db文件的时候就创建db文件
	db,err := gorm.Open("sqlite3","test.db")
	if err!=nil{
		panic("failed to connect database")
	}
	defer db.Close()

	//映射（实体类与表结构的映射）
	db.AutoMigrate(&Product{})
	fmt.Println("migrate complete")
	//创建行记录
	db.Create(&Product{Code:"L1111",Price:1000})
	fmt.Println("create complete")
	//查找记录
	var product Product
	db.First(&product,1)
	db.First(&product,"code=?","L1111")
	fmt.Println("query first record :",product.Price)
	//更新记录
	db.Model(&product).Update("Price",2000)
	fmt.Println("update first record complete :",product.Price)
	//删除记录
	db.Delete(&product)
	fmt.Println("delete record complete")
}