package go_gorm

import (
	"testing"
	"time"
)

var users []User
var users1 []User
var users2 []User
var count int64

// 查询
func Test_select_01(t *testing.T) {
	var user User
	var users []User

	// 获取第一条记录，按主键排序
	db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，不指定排序
	db.Take(&user)
	//// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录，按主键排序
	db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取所有的记录
	db.Find(&users)
	//// SELECT * FROM users;

	// 通过主键进行查询 (仅适用于主键是数字类型)
	db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
}

// Where
func Test_where_01(t *testing.T) {
	var user User
	var users []User
	// 获取第一条匹配的记录
	db.Where("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM User WHERE name = 'jinzhu' limit 1;

	// 获取所有匹配的记录
	db.Where("name = ?", "jinzhu").Find(&users)
	//// SELECT * FROM User WHERE name = 'jinzhu';

	// <>
	db.Where("name <> ?", "jinzhu").Find(&users)

	// IN
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time
	db.Where("updated_at > ?", time.Now().Add(-time.Minute*5)).Find(&users)

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", time.Now().Add(-time.Minute*5), time.Now()).Find(&users)
}

// Struct & Map
// 当通过struct进行查询的时候，GORM 将会查询这些字段的非零值， 意味着你的字段包含 0， ''， false 或者其他 零值, 将不会出现在查询语句中
// 可以将字段变为指针类型，这样就可以查询默认值(例如Age字段)
func Test_type_01(t *testing.T) {
	var user User
	var users []User
	// Struct
	var age int = 20
	db.Where(&User{Name: "jinzhu", Age: &age}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 多主键 slice 查询
	db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);
}

//  Not
func Test_not_01(t *testing.T) {
	var user User
	var users []User
	db.Not("name", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

	// 不包含
	db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	//不在主键 slice 中
	db.Not([]int64{1, 2, 3}).First(&user)
	//// SELECT * FROM users WHERE id NOT IN (1,2,3);

	db.Not([]int64{}).First(&user)
	//// SELECT * FROM users;

	// 原生 SQL
	db.Not("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE NOT(name = "jinzhu");

	// Struct
	db.Not(User{Name: "jinzhu"}).First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu";
}

// Or
func Test_or_01(t *testing.T) {
	var users []User
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';
}

// 行内条件查询
func Test_row_01(t *testing.T) {
	var users []User
	var user User
	// 通过主键进行查询 (仅适用于主键是数字类型)
	db.First(&user, 23)
	//// SELECT * FROM users WHERE id = 23 LIMIT 1;
	// 非数字类型的主键查询
	db.First(&user, "id = ?", "string_primary_key")
	//// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;

	// 原生 SQL
	db.Find(&user, "name = ?", "jinzhu")
	//// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	var age int = 20
	db.Find(&users, User{Age: &age})
	//// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	//// SELECT * FROM users WHERE age = 20;
}

// 额外的查询选项
func Test_extra_01(t *testing.T) {
	var user User
	// 为查询 SQL 添加额外的选项
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;
}

// 子查询
func Test_gaoji_01(t *testing.T) {
	var orders Order
	db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
}

// 指定字段
func Test_field_01(t *testing.T) {
	var users []User
	db.Select("name, age").Find(&users)
	//// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	//// SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	//// SELECT COALESCE(age,'42') FROM users;
}

// Order
func Test_order_01(t *testing.T) {
	var users []User
	db.Order("age desc, name").Find(&users)
	//// SELECT * FROM users ORDER BY age desc, name;

	// 多个排序条件
	db.Order("age desc").Order("name").Find(&users)
	//// SELECT * FROM users ORDER BY age desc, name;

	var users1 []User
	var users2 []User
	// 重新排序
	db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
	//// SELECT * FROM users ORDER BY age desc; (users1)
	//// SELECT * FROM users ORDER BY age; (users2)
}

// Limit
func Test_limit_01(t *testing.T) {
	var users []User
	var users1 []User
	var users2 []User
	db.Limit(3).Find(&users)
	//// SELECT * FROM users LIMIT 3;

	// 用 -1 取消 LIMIT 限制条件
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	//// SELECT * FROM users LIMIT 10; (users1)
	//// SELECT * FROM users; (users2)
}

// Offset
func Test_offset(t *testing.T) {
	db.Offset(3).Find(&users)
	//// SELECT * FROM users OFFSET 3;

	// 用 -1 取消 OFFSET 限制条件
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	//// SELECT * FROM users OFFSET 10; (users1)
	//// SELECT * FROM users; (users2)
}

// Count
func Test_Count_01(t *testing.T) {
	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
	//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

	db.Table("deleted_users").Count(&count)
	//// SELECT count(*) FROM deleted_users;
}

// Group 和 Having
func Test_Group_01(t *testing.T) {
	rows, _ := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
	for rows.Next() {
		//
	}

	rows, _ = db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	for rows.Next() {
		//
	}

	type Result struct {
		Date  time.Time
		Total int64
	}
	var results []interface{}
	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

}

//  Joins

func Test_Joins_01(t *testing.T) {
	rows, _ := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	for rows.Next() {
		//
	}

	var results []interface{}
	db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	// 多个关联查询
	var user User
	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
}

// Scan
func Test_Scan_01(t *testing.T) {
	type Result struct {
		Name string
		Age  int
	}

	var result Result
	db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result) // 将 Scan 查询结果放入另一个结构体中。

	// Raw SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
}
