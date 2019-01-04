package main

import (
	"log"

	"github.com/bangwork/ones-ai-docker/utils/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UUID() string { return uuid.V4Compressed()[0:8] }

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "bang:bang@tcp(localhost:3306)/bang?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	//for i := 3; i < 5000; i++ {
	//	uuid := UUID()
	//	fmt.Println(uuid)
	//	err = db.Exec("INSERT INTO page_tree(`uuid`,`team_uuid`,`space_uuid`,`owner_uuid`,`title`,`parent_uuid`,`status`,`position`,`create_time`,`updated_time`) VALUES(?,'3MvcLbT7','F5ebej19','4HcfEm3h','1111','',1,?,1540537492,1540537492);", uuid, i).Error
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//}
	uuid1:=UUID()
	sql11:="insert into `permission_rule`(`uuid`,`team_uuid`,`context_type`,`context_param_1`,`user_domain_type`,`user_domain_param`,`permission`,`create_time`,`status`)"
	sql12:=	" values(?,'3MvcLbT7',1102,'F5ebej19',3,'',2002,1540537492,1)"
	sql1:=sql11+sql12
	err = db.Exec(sql1,uuid1).Error
	if err!=nil{
		log.Fatal(err)
	}

	uuid2:=UUID()
	sql21:="insert into `permission_rule`(`uuid`,`team_uuid`,`context_type`,`context_param_1`,`user_domain_type`,`user_domain_param`,`permission`,`create_time`,`status`)"
	sql22:=	" values(?,'3MvcLbT7',1102,'F5ebej19',3,'4HcfEm3h',2001,1540537492,1)"
	sql2:=sql21+sql22
	err = db.Exec(sql2,uuid2).Error
	if err!=nil{
		log.Fatal(err)
	}

	uuid3:=UUID()
	sql31:="insert into `permission_rule`(`uuid`,`team_uuid`,`context_type`,`context_param_1`,`user_domain_type`,`user_domain_param`,`permission`,`create_time`,`status`)"
	sql32:=	" values(?,'3MvcLbT7',1102,'F5ebej19',3,'',2003,1540537492,1)"
	sql3:=sql31+sql32
	err = db.Exec(sql3,uuid3).Error
	if err!=nil{
		log.Fatal(err)
	}
}
