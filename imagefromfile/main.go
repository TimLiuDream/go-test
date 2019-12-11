package main

import (
	"baliance.com/gooxml/common"
	"fmt"
	"log"
)

func main() {
	imgPath := "./imagefromfile/image/UDocQz4H.jpg"
	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(img)
}
