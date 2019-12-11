package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

const (
	Kilobyte = 1024
	Megabyte = 1024 * Kilobyte

	DEFAULT_MAX_WIDTH  float64 = 1920
	DEFAULT_MAX_HEIGHT float64 = 1080
)

func main() {
	src_pic_path := "/Users/tim/Downloads/RSAKuGRJ.jpg"
	compressed_pic_path := "/Users/tim/Downloads/RSAKuGRJcompress.jpg"

	if !imageCompress(
		func() (io.Reader, error) {
			return os.Open(src_pic_path)
		},
		func() (*os.File, error) {
			return os.Open(src_pic_path)
		},
		compressed_pic_path, 95, 1440, "jpeg") {
		fmt.Println("生成缩略图失败")
		return
	}
	fmt.Println("生成缩略图成功 " + compressed_pic_path)

	// fileinfo, err := os.Stat(src_pic_path)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// filesize := fileinfo.Size()
	// fmt.Println(filesize)

	// if filesize > 4*Megabyte {
	// 	fmt.Printf("is need compress?:%t", true)
	// 	err = makeThumbnail(src_pic_path, compressed_pic_path)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return
	// }
	// fmt.Printf("is need compress?:%t", false)
}

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func makeThumbnail(imagePath, savePath string) error {

	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.NearestNeighbor)

	// 需要保存的文件
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// 以jpg格式保存文件
	err = jpeg.Encode(imgfile, m, nil)
	if err != nil {
		return err
	}

	return nil
}

func imageCompress(
	getReadSizeFile func() (io.Reader, error),
	getDecodeFile func() (*os.File, error),
	to string,
	Quality,
	base int,
	format string) bool {
	/** 读取文件 */
	file_origin, err := getDecodeFile()
	defer file_origin.Close()
	if err != nil {
		fmt.Println("os.Open(file)错误")
		log.Fatal(err)
		return false
	}
	var origin image.Image
	var config image.Config
	var temp io.Reader
	/** 读取尺寸 */
	temp, err = getReadSizeFile()
	if err != nil {
		fmt.Println("os.Open(temp)")
		log.Fatal(err)
		return false
	}
	var typeImage int64
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format == "jpg" || format == "jpeg" {
		typeImage = 1
		origin, err = jpeg.Decode(file_origin)
		if err != nil {
			fmt.Println("jpeg.Decode(file_origin)")
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)")
			log.Fatal(err)
			return false
		}
		config, err = jpeg.DecodeConfig(temp)
		if err != nil {
			fmt.Println("jpeg.DecodeConfig(temp)")
			return false
		}
	} else if format == "png" {
		typeImage = 0
		origin, err = png.Decode(file_origin)
		if err != nil {
			fmt.Println("png.Decode(file_origin)")
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)")
			log.Fatal(err)
			return false
		}
		config, err = png.DecodeConfig(temp)
		if err != nil {
			fmt.Println("png.DecodeConfig(temp)")
			return false
		}
	}
	/** 做等比缩放 */
	width := uint(base) /** 基准 */
	height := uint(base * config.Height / config.Width)

	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if typeImage == 0 {
		err = png.Encode(file_out, canvas)
		if err != nil {
			fmt.Println("压缩图片失败")
			return false
		}
	} else {
		err = jpeg.Encode(file_out, canvas, &jpeg.Options{Quality})
		if err != nil {
			fmt.Println("压缩图片失败")
			return false
		}
	}

	return true
}
