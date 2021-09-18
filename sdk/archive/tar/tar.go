package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ArchiveTar struct {
	Files []struct {
		Name string
		Body []byte
	}
}

func (a *ArchiveTar) Reset() {
	a.Files = nil
}

func (a *ArchiveTar) Read(buf *bytes.Buffer) (err error) {
	// 打开tar档案以供阅读。
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	return a.reader(tr)
}

func (a *ArchiveTar) ReadFile(file string) (err error) {
	bys, err := readFile(file)
	if err != nil {
		return err
	}
	r := bytes.NewReader(bys)
	tr := tar.NewReader(r)

	return a.reader(tr)
}

func (a *ArchiveTar) reader(reader *tar.Reader) error {
	// 迭代档案中的文件。
	for {
		hdr, err := reader.Next()
		if err == io.EOF {
			// tar归档结束
			break
		}
		if err != nil {
			return err
		}

		bys := new(bytes.Buffer)

		_, err = io.Copy(bys, reader)
		if err != nil {
			return err
		}

		a.Files = append(a.Files, struct {
			Name string
			Body []byte
		}{Name: hdr.Name, Body: bys.Bytes()})
	}

	return nil
}

func (a *ArchiveTar) Write() (buf *bytes.Buffer, err error) {
	// 创建一个缓冲区来写入我们的存档。
	buf = new(bytes.Buffer)
	// 创建一个新的tar存档。
	tw := tar.NewWriter(buf)
	{
		for _, file := range a.Files {
			hdr := &tar.Header{
				Name: file.Name,
				Mode: 0600,
				Size: int64(len(file.Body)),
			}
			if err = tw.WriteHeader(hdr); err != nil {
				return
			}
			if _, err = tw.Write([]byte(file.Body)); err != nil {
				return
			}
		}
	}
	// 确保在Close时检查错误。
	if err = tw.Close(); err != nil {
		return
	}
	return
}

func (a *ArchiveTar) WriteFile(file string) error {
	buf, err := a.Write()
	if err != nil {
		return err
	}

	return writeFile(file, buf.Bytes())
}

func readFile(file string) (bys []byte, err error) {
	return ioutil.ReadFile(file)
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
