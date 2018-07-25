package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	exists, err := PathExists("./test/")
	if err != nil {
		fmt.Println("get dir error:", err)
		return
	}

	if exists {
		fmt.Println("dir is exist!")
	} else {
		fmt.Println("dir is not exist,creating now")
		err := os.Mkdir("./test/", os.ModePerm)
		if err != nil {
			fmt.Println("mkdir failed")
		} else {
			fmt.Println("mkdir success!")
		}
	}
}
