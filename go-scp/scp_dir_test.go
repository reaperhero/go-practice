package go_scp

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"path/filepath"
	"testing"
	"time"
)

func ScpDir(srcClient, destClient *ssh.Client, sourceDir, destDir string) error {
	srcSftpClient, err := sftp.NewClient(srcClient)
	if err != nil {
		return err
	}
	defer srcSftpClient.Close()

	destSftpClient, err := sftp.NewClient(destClient)
	if err != nil {
		return err
	}
	defer destSftpClient.Close()
	err = destSftpClient.MkdirAll(destDir)
	if err != nil {
		return err
	}

	fileInfos, err := srcSftpClient.ReadDir(sourceDir)
	for _, info := range fileInfos {
		if info.IsDir() {
			srcDirName := filepath.Join(sourceDir, info.Name())
			destDirName := filepath.Join(destDir, info.Name())
			err := destSftpClient.MkdirAll(destDirName)
			if err != nil {
				return err
			}
			err = ScpDir(srcClient, destClient, srcDirName, destDirName)
			if err != nil {
				return err
			}
			continue
		}

		old, err := srcSftpClient.Open(filepath.Join(sourceDir, info.Name()))
		if err != nil {
			return err
		}
		newFile, err := destSftpClient.Create(filepath.Join(destDir, info.Name()))
		if err != nil {
			return err
		}
		_, err = io.Copy(newFile, old)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestSCP(t *testing.T) {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("Abc!@#135"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}

	srcClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", "172.16.82.19", 22), config)
	if err != nil {
		fmt.Println(err)
		return
	}
	dstClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", "172.16.82.234", 22), config)
	if err != nil {
		fmt.Println(err)
		return
	}
	ScpDir(srcClient, dstClient, "/root/logs", "/root/sadsa")
}
