package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test_insert_01(t *testing.T) {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	defer Db.Close()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func TestInsterTypevalue(t *testing.T) {
	a1 := []string{"a", "1", "2"} // [a 1 2]
	a2 := []int{1, 2, 3}          // [1 2 3]
	Db.Exec("insert into person(name) values(?)", fmt.Sprint(a1))
	Db.Exec("insert into person(name) values(?)", fmt.Sprint(a2))

}
