package gziputil

//解压

import (
	"archive/tar"
	"compress/gzip"
	"github.com/gogf/gf/os/gfile"
	"io"
	"os"
	"path"
)

func Decompress(srcFilePath string, destDirPath string) error {
	os.Mkdir(destDirPath, os.ModePerm)

	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if hdr.Typeflag != tar.TypeDir {
			dir := destDirPath + "/" + path.Dir(hdr.Name)
			if !gfile.Exists(dir) {
				err := os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return err
				}
			}
			fw, _ := os.OpenFile(destDirPath+"/"+hdr.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
