package file

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// 文件末尾插入内容
func appendStringInFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("文件打开失败: %v", err)
	}
	defer file.Close()
	// 查找文件末尾的偏移量
	n, _ := file.Seek(0, io.SeekEnd)
	// 从末尾的偏移量开始写入内容
	_, err = file.WriteAt([]byte("\n"+content), n)
	if err != nil {
		logrus.Fatalf("文件写入失败: %v", err)
	}
}


// 文件指定行后插入内容
func searchAndWriteInfConst(fileName, content, infoName string) error {
	lineBytes, err := ioutil.ReadFile(fileName)
	var lines []string
	if err != nil {
		fmt.Println(err)
	} else {
		contents := string(lineBytes)
		lines = strings.Split(contents, "\n")
	}
	var newLines []string

	for _, line := range lines {
		newLines = append(newLines, line)
		isIn, err := regexp.MatchString(infoName, line)
		if err != nil {
			logrus.Printf("匹配常量类里的interface报错 :%v", err)
			continue
		}
		if isIn {
			newLines = append(newLines, content)
		}
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	defer file.Close()
	_, err = file.WriteString(strings.Join(newLines, "\n"))
	if err != nil {
		return err
	}
	return nil
}
