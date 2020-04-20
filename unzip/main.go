package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DeCompress(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		if strings.HasPrefix(file.Name, "__MACOSX/._") {
			continue
		}
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := DeCompress("/Users/tim/go/src/github.com/timliudream/golangtraining/unzip/manual-测试pandoc导入.docx.zip", "/Users/tim/go/src/github.com/timliudream/golangtraining/unzip/test/")
	if err != nil {
		fmt.Println(err)
	}
}
