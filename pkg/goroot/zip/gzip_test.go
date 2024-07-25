package zip

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// When N < 10  ==> underlying error
// When N >= 10 ==> gzip: invalid header
var N = 1

func TestGzipName(t *testing.T) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("YourDataHere")); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(str)
	data, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(data)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	fmt.Println(string(s))
}

func TestGzipFile(t *testing.T) {
	fi, _ := os.Open("/Users/edy/GolandProjects/easymatrix/matrix/install_agentx.sh")
	content, _ := io.ReadAll(fi)
	fi.Close()

	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)
	if _, err := w.Write(content); err != nil {
		logrus.Errorf("Write: %v", err)
	}
	if err := w.Flush(); err != nil {
		logrus.Errorf("Write Flush: %v", err)
	}
	if err := w.Close(); err != nil {
		logrus.Errorf("Writer.Close: %v", err)
	}

	reader, err := gzip.NewReader(bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		logrus.Errorf("Writer.Close: %v", err)
	}
	reader.Close()
}
