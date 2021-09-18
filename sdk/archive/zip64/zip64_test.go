package zip64

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

/*
	https://cloud.tencent.com/developer/section/1140386
*/

func Test_SDK_Zip64(t *testing.T) {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)

	// 将一些文件添加到存档中。
	buf, err := (&ArchiveZip64{Files: []struct {
		Name string
		Body []byte
	}{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling licence.\nWrite more examples.")},
	}}).Write()

	if err != nil {
		log.Fatalln(err)
	}

	zr := &ArchiveZip64{}
	err = (zr).Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	for _, files := range zr.Files {
		fmt.Println("Contents of :", files.Name, "\n", string(files.Body))
	}
}

func Test_SDK_Zip64_rwfile(t *testing.T) {

	// 将一些文件添加到存档中。
	err := (&ArchiveZip64{Files: []struct {
		Name string
		Body []byte
	}{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling licence.\nWrite more examples.")},
	}}).WriteFile("testdata/zip64.zip")

	if err != nil {
		log.Fatalln(err)
	}

	zr := &ArchiveZip64{}
	err = (zr).ReadFile("testdata/zip64.zip")
	if err != nil {
		log.Fatalln(err)
	}
	for _, files := range zr.Files {
		fmt.Println("Contents of :", files.Name, "\n", string(files.Body))
	}
}
