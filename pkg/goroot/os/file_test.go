package os

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// master-configmap.yaml mysql-sts.yaml        slave-configmap.yaml

func TestFileRead(t *testing.T) {

	filename, err := os.OpenFile("/Users/chenqiangjun/github/k8s-application/infrastructure/statefulset/mysql/master-configmap.yaml", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//filename.Truncate(0) // 清空文件
	all, err := io.ReadAll(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(all))
}
