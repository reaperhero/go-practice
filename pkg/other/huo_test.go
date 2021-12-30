package other

import (
"testing"
"fmt"
)


// & 其功能是参与运算的两数各对应的二进位相与。
// | 其功能是参与运算的两数各对应的二进位相或
// ^ 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1


//位运算符对整数在内存中的二进制位进行操作。
//p	q  p&q p|q p^q
//0	0	0	0	0
//0	1	0	1	1
//1	1	1	1	0
//1	0	0	1	1

func TestHu(t *testing.T)  {
	str := ""
	for i:=0;i<13;i++ {
		str += "../"
	}
	fmt.Println(str)
}
