package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessKey := os.Getenv("QINIUACCESSKEY")
	secretKey := os.Getenv("QINIUSECRETKEY")
	bucket := os.Getenv("QINIUBUCKETNAME")
	storgeURL := os.Getenv("STORGEURL")

	fmt.Println("Please enter the local file path: ")
	localFile := ""
	fmt.Scanln(&localFile)

	localFile = strings.TrimSpace(localFile)
	keys := strings.Split(localFile, "/")
	key := keys[len(keys)-1]

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(storgeURL + key)
}
