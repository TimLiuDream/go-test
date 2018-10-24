package main

import (
	"log"
	"os"
	_ "GolangTraining/simple1/matchers"
	"GolangTraining/simple1/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//使用特定的项做搜索
	search.Run("Willa")
}
