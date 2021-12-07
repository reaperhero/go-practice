package file

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*在文件写入内容，没有文件则重新创建*/
func TestWriteBybufio(t *testing.T) {
	filePath := "./1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) // |os.O_TRUNC 会清空文件
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	defer file.Close()
	str := "hello world\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*因为writer是带缓存的，需要通过flush到磁盘*/
	writer.Flush()
}