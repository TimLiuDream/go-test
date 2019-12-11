package main

import (
	"archive/zip"
	"fmt"
	"log"
)

const (
	FilePath = "samples/sample1.zip"
)

func main() {
	r, err := zip.OpenReader(FilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("contents of %s:\n", f.Name)
		// rc, err := f.Open()
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// _, err = io.CopyN(os.Stdout, rc, 68)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// rc.Close()
		// fmt.Println()
	}
}
