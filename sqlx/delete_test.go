package sqlx

import (
	"fmt"
	"testing"
)

func Test_sqlx_04(t *testing.T)  {
	res, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row,err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}

	fmt.Println("delete succ: ",row)
}