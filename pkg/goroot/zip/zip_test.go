package zip

import (
	"archive/zip"
	"bytes"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLocalZip(t *testing.T) {

	b, _ := os.OpenFile("1.zip", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644) // 1.zip 1 22  todo.txt
	// Create a new zip archive.
	w := zip.NewWriter(b)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"11/readme.txt", "This archive contains some text files."},
		{"22/gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			logrus.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			logrus.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestName(t *testing.T) {
	var (
		buf = new(bytes.Buffer)
		w   = zip.NewWriter(buf)
	)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"11/readme.txt", "This archive contains some text files."},
		{"22/gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			logrus.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			logrus.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		logrus.Fatal(err)
	}

	//ctx.Header("Content-Disposition", "attachment; filename=dashbroad.zip")
	//ctx.Header("Content-Type", "application/zip")
	//_, err = buf.WriteTo(ctx.ResponseWriter())
}
