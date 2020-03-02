package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func extractAndWrite(file *zip.File) error {
	rc, err := file.Open()

	if err != nil {
		return err
	}

	defer rc.Close()

	path, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), file.Name)

	if err != nil {
		return err
	}
	if file.FileInfo().IsDir() {
		os.MkdirAll(path, 0755)

		return nil
	}

	os.MkdirAll(filepath.Dir(path), 0755)

	output, err := os.Create(path)

	if err != nil {
		return err
	}
	defer output.Close()

	if _, err = io.Copy(output, rc); err != nil {
		return err
	}

	return nil
}

func Unzip(src string) error {
	reader, err := zip.OpenReader(src)

	if err != nil {
		return err
	}

	defer reader.Close()

	for _, file := range reader.File {
		if err := extractAndWrite(file); err != nil {
			return err
		}
	}

	return nil
}
