package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

//读写文件
/*在已存在文件清空原有内容进行追加*/
func TestOpenFile(t *testing.T) {
	filePath := "./1.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	/*关闭文件流*/
	defer file.Close()
	/*读取*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	/*写入文件*/
	str := "hello FCC您好！！！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*因为writer是带缓存的，需要通过flush到磁盘*/
	writer.Flush()
}

/*将文件1的内容拷贝到文件2*/
func TestFileCopy01(t *testing.T) {
	file1Path := "./1.txt"
	file2Path := "./2.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
	}
}

/*文件的拷贝*/
func TestFileCopy02(t *testing.T) {
	srcFile, err := os.Open("srcFileName.txt")
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile("dstFileName.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	writeN, err := io.Copy(writer, reader)
	fmt.Println(writeN, err)
}

/*判断文件以及目录是否存在*/
func TestPathExists(t *testing.T) {
	_, err := os.Stat("./1.txt")
	if err == nil {
		fmt.Println("当前文件存在！")
	}
	if os.IsNotExist(err) {
		fmt.Println("当前文件不存在！")
	}
}

/*统计文件的字符个数*/
func TestCharCount(t *testing.T) {
	type CharCount struct {
		/*英文的个数*/
		ChCount int
		/*数字的个数*/
		NumCount int
		/*空格的个数*/
		SpaceCount int
		/*其他字符的个数*/
		OtherCount int
	}
	fileName := "./1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 强制执行下一行case代码,不需要再判断
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为：%v 数字的个数为：%v 空格的个数为：%v 其他字符的个数为：%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
