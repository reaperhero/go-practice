package go_gorm

import (
	"testing"
	"time"
)

func Test_create_record(t *testing.T)  {
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	db.NewRecord(user) // => 返回 `true` ，因为主键为空

	db.Create(&user)

	db.NewRecord(user) // => 在 `user` 之后创建返回 `false`
}