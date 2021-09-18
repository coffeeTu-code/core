package zip64

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ArchiveZip64 struct {
	Files []struct {
		Name string
		Body []byte
	}
}

func (a *ArchiveZip64) Reset() {
	a.Files = nil
}

func (a *ArchiveZip64) Read(buf *bytes.Buffer) (err error) {
	// 打开一个zip文件供阅读。
	r, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(len(buf.Bytes())))
	//r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		return err
	}

	return a.reader(r)
}

func (a *ArchiveZip64) ReadFile(file string) (err error) {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer r.Close()

	return a.reader(&r.Reader)
}

func (a *ArchiveZip64) reader(reader *zip.Reader) (err error) {
	// 迭代档案中的文件。
	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		bys := new(bytes.Buffer)

		_, err = io.Copy(bys, rc)
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}
		rc.Close()

		a.Files = append(a.Files, struct {
			Name string
			Body []byte
		}{Name: f.Name, Body: bys.Bytes()})
	}
	return
}

func (a *ArchiveZip64) Write() (buf *bytes.Buffer, err error) {
	// 创建一个缓冲区来写入我们的存档。
	buf = new(bytes.Buffer)
	// 创建一个新的tar存档。
	zw := zip.NewWriter(buf)
	{
		for _, file := range a.Files {
			f, err := zw.Create(file.Name)
			if err != nil {
				return nil, err
			}
			_, err = f.Write([]byte(file.Body))
			if err != nil {
				return nil, err
			}
		}
	}
	// 确保在Close时检查错误。
	if err = zw.Close(); err != nil {
		return
	}
	return
}

func (a *ArchiveZip64) WriteFile(file string) (err error) {
	buf, err := a.Write()
	if err != nil {
		return err
	}

	return writeFile(file,buf.Bytes())
}

func writeFile(file string, bys []byte) error {
	if err := mkdir(file); err != nil {
		return err
	}
	return ioutil.WriteFile(file, bys, 0600)
}

func mkdir(path string) error {
	dir, _ := filepath.Split(path)
	if len(dir) > 0 {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}
