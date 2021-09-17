package zip64

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

/*
	https://cloud.tencent.com/developer/section/1140386
*/

func Test_SDK_Zip64(t *testing.T) {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)

	// 创建一个新的zip存档。
	w := zip.NewWriter(buf)

	// 将一些文件添加到存档中。将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 确保在Close时检查错误。
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 打开一个zip文件供阅读。
	r, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(len(buf.Bytes())))
	//r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	//defer r.Close()

	// 遍历存档中的文件，
	// 打印他们的一些内容。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
