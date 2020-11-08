package curd

import (
	"testing"
)

// 删除记录
// 警告：当删除一条记录的时候，你需要确定这条记录的主键有值，GORM会使用主键来删除这条记录。如果主键字段为空，GORM会删除模型中所有的记录。
func Test_delete_01(t *testing.T) {
	var email Email
	// 删除一条存在的记录
	db.Delete(&email)
	//// DELETE from emails where id=10;

	// 为删除 SQL 语句添加额外选项
	db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&email)
	//// DELETE from emails where id=10 OPTION (OPTIMIZE FOR UNKNOWN);
}

// 批量删除
func Test_batch_delete_01(t *testing.T) {
	db.Where("email LIKE ?", "%jinzhu%").Delete(Email{})
	//// DELETE from emails where email LIKE "%jinzhu%";

	db.Delete(Email{}, "email LIKE ?", "%jinzhu%")
	//// DELETE from emails where email LIKE "%jinzhu%";
}

// 软删除
func Test_soft_delete(t *testing.T) {
	db.Delete(&user)
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// 在查询记录时，软删除记录会被忽略
	db.Where("age = 20").Find(&user)
	//// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

	// 使用 Unscoped 方法查找软删除记录
	db.Unscoped().Where("age = 20").Find(&users)
	//// SELECT * FROM users WHERE age = 20;

	// 使用 Unscoped 方法永久删除记录
	var order Order
	db.Unscoped().Delete(&order)
	//// DELETE FROM orders WHERE id=10;
}
