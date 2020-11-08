package curd

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {
	db, _ = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbuser?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	db.SingularTable(true) // 禁用表名复数形式
}
