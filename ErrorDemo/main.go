package main

import "errors"
import "fmt"

func main() {
	err := errors.New("emit macho dwarf: elf handle corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
