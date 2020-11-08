package curd

import (
	"github.com/jinzhu/gorm"
	"go-example-demo/go-gorm"
	"testing"
)

var (
	user User
)

// 更新所有字段
// Save 方法在执行 SQL 更新操作时将包含所有字段，即使这些字段没有被修改。
func Test_save_01(t *testing.T) {
	db.First(&user)
	age := 100
	user.Name = "jinzhu 2"
	user.Age = &age
	db.Save(&user) // Save 方法在执行 SQL 更新操作时将包含所有字段，即使这些字段没有被修改

	//// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
}

// 更新已更改的字段
// 如果你只想更新已经修改了的字段，可以使用 Update，Updates 方法
func Test_filed_01(t *testing.T) {
	// 如果单个属性被更改了，更新它
	db.Model(&user).Update("name", "hello")
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用组合条件更新单个属性
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// 使用 `map` 更新多个属性，只会更新那些被更改了的字段
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	//// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用 `struct` 更新多个属性，只会更新那些被修改了的和非空的字段
	var age int = 18
	db.Model(&user).Updates(User{Name: "hello", Age: &age})
	//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

}

// 更新选中的字段
// 如果你在执行更新操作时只想更新或者忽略某些字段，可以使用 Select，Omit方法。
func Test_Select_02(t *testing.T) {
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	//// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
}

// 更新列钩子方法
// 上面的更新操作更新时会执行模型的 BeforeUpdate 和 AfterUpdate 方法，来更新 UpdatedAt 时间戳，并且保存他的 关联。如果你不想执行这些操作，可以使用 UpdateColumn，UpdateColumns 方法
func Test_Column_01(t *testing.T) {
	// Update single attribute, similar with `Update`
	db.Model(&user).UpdateColumn("name", "hello")
	//// UPDATE users SET name='hello' WHERE id = 111;

	// Update multiple attributes, similar with `Updates`
	var age int = 18
	db.Model(&user).UpdateColumns(User{Name: "hello", Age: &age})
	//// UPDATE users SET name='hello', age=18 WHERE id = 111;
}

// 批量更新
// 批量更新时，钩子函数不会执行
func Test_batch_01(t *testing.T) {
	db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	// 使用结构体更新将只适用于非零值，或者使用 map[string]interface{}
	var age int = 18
	db.Model(User{}).Updates(User{Name: "hello", Age: &age})
	//// UPDATE users SET name='hello', age=18;

	// 使用 `RowsAffected` 获取更新影响的记录数
	_ = db.Model(User{}).Updates(User{Name: "hello", Age: &age}).RowsAffected

}

// 带有表达式的 SQL 更新
func Test_Expr_01(t *testing.T) {
	var product Product
	db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

	db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

	db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';

	db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;
}
