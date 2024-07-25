package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"io"
	"os"
)

const address = "master:9000"

func main() {

	client, err := hdfs.New(address)
	if err != nil {
		panic(err)
	}

	// 创建目录
	{
		path := "/testdir"
		err = client.MkdirAll(path, 0777) // 创建testdir目录
		if err != nil {
			panic(err)
		}

		fmt.Printf("Created directory: %s\n", path)
	}

	// 上传文件
	{
		localPath := "./file.txt"
		hdfsPath := "/testdir/file.txt"

		// 打开本地文件
		localFile, err := os.Open(localPath)
		if err != nil {
			panic(err)
		}
		defer localFile.Close()

		// 创建HDFS文件
		hdfsFile, err := client.Create(hdfsPath)
		if err != nil {
			panic(err)
		}
		defer hdfsFile.Close()

		// 将本地文件复制到HDFS
		_, err = io.Copy(hdfsFile, localFile)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Uploaded file: %s\n", hdfsPath)
	}

	// 下载文件
	{
		hdfsPath := "/test.txt"
		localPath := "/home/ubuntu/workspace/hadoop/test.txt"

		hdfsFile, err := client.Open(hdfsPath)
		if err != nil {
			panic(err)
		}
		defer hdfsFile.Close()

		localFile, err := os.Create(localPath)
		if err != nil {
			panic(err)
		}
		defer localFile.Close()

		_, err = io.Copy(localFile, hdfsFile)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded file: %s\n", localPath)
	}

	// 查看文件列表
	{
		hdfsPath := "/testdir"

		files, err := client.ReadDir(hdfsPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Files in %s:\n", hdfsPath)
		for _, file := range files {
			fmt.Printf("%s (size: %d)\n", file.Name(), file.Size())
		}
	}

	// 删除文件
	{
		hdfsPath := "/testdir"

		err = client.Remove(hdfsPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Deleted directory: %s\n", hdfsPath)
	}
}
