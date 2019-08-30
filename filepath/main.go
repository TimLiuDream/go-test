package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	filePaths := make([]string, 0)
	rootPath := "/Users/tim/go/src/github.com/bangwork/bang-api/migration/F9008"
	err := filepath.Walk(rootPath, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", p, err)
			return err
		}
		if !info.IsDir() && info.Name() != "bindata.go" {
			filePaths = append(filePaths, p)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}
