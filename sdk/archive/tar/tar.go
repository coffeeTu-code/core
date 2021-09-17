package tar

import (
	"archive/tar"
	"bytes"
	"io"
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

	// 迭代档案中的文件。
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// tar归档结束
			break
		}
		if err != nil {
			return err
		}

		bys := new(bytes.Buffer)

		_, err = io.Copy(bys, tr)
		if err != nil {
			return err
		}

		a.Files = append(a.Files, struct {
			Name string
			Body []byte
		}{Name: hdr.Name, Body: bys.Bytes()})
	}
	return
}

func (a *ArchiveTar) ReadFile() {

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

func (a *ArchiveTar) WriteFile() {

}
