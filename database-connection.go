package main

import (
	"fmt"
 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
// MYSQL DATABASE [ LOCALHOST ] ID --------------->
var urlDSN = "WJ28@krhps:WJ28@krhps@tcp(localhost:3306)/database?parseTime=true"
var err error

func DataMigration() {

	Database, err =  gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		
		fmt.Print(err.Error())
		panic("connection faild :(")
	}
	 Database.AutoMigrate(&Employee{})
	
	 


}



