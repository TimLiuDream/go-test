package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("/Users/tim/Documents/base64.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := strings.Replace(string(bytes), "\n", "", -1)
	result, err = stripMime(result)
	if err != nil {
		log.Fatalln(err)
	}
	imgPath := base64ToImage(result)
	fmt.Println(imgPath)
}

func base64ToImage(base64Str string) (imgPath string) {
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}
	imgPath = fmt.Sprintf("./base64toimage/%s", UUID.String()+".jpg")
	ddd, _ := base64.RawStdEncoding.DecodeString(base64Str)
	err = ioutil.WriteFile(imgPath, ddd, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func stripMime(combined string) (string, error) {
	re := regexp.MustCompile("data:(.*);base64,(.*)")
	parts := re.FindStringSubmatch(combined)

	if len(parts) < 3 {
		return "", errors.New("Invalid base64 input")
	}

	data := parts[2]
	return data, nil
}
