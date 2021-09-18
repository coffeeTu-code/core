package tar

import (
	"fmt"
	"log"
	"testing"
)

/*
	https://cloud.tencent.com/developer/section/1140383
*/

func Test_SDK_Tar(t *testing.T) {

	// 将一些文件添加到存档中。
	buf, err := (&ArchiveTar{Files: []struct {
		Name string
		Body []byte
	}{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}}).Write()

	if err != nil {
		log.Fatalln(err)
	}

	tr := &ArchiveTar{}
	err = (tr).Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	for _, files := range tr.Files {
		fmt.Println("Contents of :", files.Name, "\n", string(files.Body))
	}
}

func Test_SDK_Tar_rwfile(t *testing.T) {

	// 将一些文件添加到存档中。
	err := (&ArchiveTar{Files: []struct {
		Name string
		Body []byte
	}{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}}).WriteFile("testdata/tar.tar")

	if err != nil {
		log.Fatalln(err)
	}

	tr := &ArchiveTar{}
	err = (tr).ReadFile("testdata/tar.tar")
	if err != nil {
		log.Fatalln(err)
	}
	for _, files := range tr.Files {
		fmt.Println("Contents of :", files.Name, "\n", string(files.Body))
	}
}
