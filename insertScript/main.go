package main

import (
	"fmt"
	"log"

	"github.com/bangwork/ones-ai-docker/utils/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UUID() string { return uuid.V4Compressed()[0:8] }

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "bang:bang@tcp(localhost:3306)/wiki?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := 3; i < 10000; i++ {
		uuid := UUID()
		fmt.Println(uuid)
		err = db.Exec("INSERT INTO page_tree(`uuid`,`team_uuid`,`space_uuid`,`owner_uuid`,`title`,`parent_uuid`,`status`,`position`,`create_time`,`updated_time`) VALUES(?,'3MvcLbT7','F5ebejM6','4HcfEm3h','1111','',1,?,1540537492,1540537492);", uuid, i).Error
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
