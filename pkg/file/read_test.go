package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadByOs(t *testing.T) {
	file, err := os.Open("./1.jpg")
	if err != nil {
		fmt.Println("open file err", err)
	}
	fmt.Printf("file=%v", file)
	err1 := file.Close()
	if err1 != nil {
		fmt.Println("close file err = ", err1)
	}
}

/*缓冲式读取文件*/
func TestReadBybufio(t *testing.T) {
	file, err := os.Open("./1.jpg")
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束!")
}

/*ioutil读取文件*/
func TestReadByIoutil(t *testing.T) {
	file := "./1.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}
