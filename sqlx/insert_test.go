package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
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

//

func TestBatchInsert(t *testing.T) {
	type ggg struct {
		Name       string    `db:"name"`
		UpdateTime time.Time `db:"update_time"`
	}
	gs := []ggg{
		ggg{
			Name:       "1",
			UpdateTime: time.Now(),
		},
		ggg{
			Name:       "2",
			UpdateTime: time.Now(),
		},
	}
	result,err := Db.NamedExec(fmt.Sprintf("insert into actor (name,update_time) VALUES ('%s',:update_time)","asds"),gs)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(result.RowsAffected())
}
