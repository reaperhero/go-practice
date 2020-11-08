package sqlx

import (
	"fmt"
	"testing"
)

func Test_insert_02(t *testing.T) {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}
