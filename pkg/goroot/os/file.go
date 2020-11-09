package os

import (
	"os"
)

var (
	filename, _ = os.OpenFile("/tmp/register.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
)

func main() {
	filename.Truncate(0) // 清空文件
}
