package go_scp

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
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

	return scpDir(srcSftpClient, destSftpClient, sourceDir, destDir)
}

func scpDir(srcClient, dstClient *sftp.Client, sourceDir, destDir string) error {
	fileInfos, err := srcClient.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	err = dstClient.MkdirAll(destDir)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			srcDirName := filepath.Join(sourceDir, info.Name())
			destDirName := filepath.Join(destDir, info.Name())
			err = scpDir(srcClient, dstClient, srcDirName, destDirName)
			if err != nil {
				return err
			}
			continue
		}
		old, err := srcClient.Open(filepath.Join(sourceDir, info.Name()))
		if err != nil {
			return err
		}
		newFile, err := dstClient.Create(filepath.Join(destDir, info.Name()))
		if err != nil {
			return err
		}
		logrus.Printf("copy %s to %s start", old.Name(), newFile.Name())
		_, err = io.Copy(newFile, old)
		if err != nil {
			logrus.Errorf("copy %s to %s %v", old.Name(), newFile.Name(), err)
			return err
		}
		logrus.Infof("copy %s to %s", old.Name(), newFile.Name())

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

func SshExecCmd(sshClient *ssh.Client, cmd string) (stdout, stderr string, err error) {
	session, err := sshClient.NewSession()
	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	if err != nil {
		return "", "New ssh session failed", err
	}
	defer session.Close()
	err = session.Run(cmd)
	// 没有报错的时候stdoutBuf stderrBuf 内容一样
	if err != nil {
		return stdoutBuf.String(), stderrBuf.String(), err
	}
	return stdoutBuf.String(), "", nil
}

func TestCmd(t *testing.T) {
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
	stdOut, stdErr, err := SshExecCmd(srcClient, "sh /tmp/sss.sh")
	if err != nil {
		return
	}
	fmt.Println(stdOut, stdErr, err)
}
