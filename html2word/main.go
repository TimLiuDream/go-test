package main

import (
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bangwork/ones-ai-api-common/utils/uuid"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	sourcePath := "./html2word/test.html"
	targetPath := "./html2word/test.docx"
	file, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}
	htmlDoc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln(err)
	}

	doc := document.New()
	paragraph := doc.AddParagraph()
	run := paragraph.AddRun()
	htmlDoc.Find("body").Children().Each(func(i int, selection *goquery.Selection) {
		for _, node := range selection.Nodes {
			parseElement(node, doc, run)
		}
	})

	err = doc.SaveToFile(targetPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseElement(node *html.Node, doc *document.Document, run document.Run) {
	if node.Type == html.TextNode {
		if strings.TrimSpace(node.Data) != "" {
			// 直接把文本写进去文件
			run.AddText(node.Data)
		}
	} else if node.Type == html.ElementNode {
		tag := node.DataAtom.String()
		if tag == "p" || tag == "li" {
			// 直接把换行写进文件中
			run.AddText("\n")
		} else if tag == "figure" {
			// 把图片写进文件
			parseImg(node, doc, run)
		}
	}
	if node.FirstChild != nil {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			parseElement(c, doc, run)
		}
	}
}

func parseImg(node *html.Node, doc *document.Document, run document.Run) {
	if node.FirstChild != nil {
		c := node.FirstChild.NextSibling.FirstChild
		attr := c.Attr[1]
		base64Str := strings.Replace(attr.Val, "\n", "", -1)
		base64Str, err := stripMime(base64Str)
		if err != nil {
			log.Fatalln(err)
		}
		imgPath := base2img(base64Str)
		// 自己构造image属性，先查出图片的宽高，再设置image的属性
		img, err := common.ImageFromFile(imgPath)
		if err != nil {
			log.Fatalln(err)
		}
		imgRef, err := doc.AddImage(img)
		if err != nil {
			log.Fatalln(err)
		}
		anchored, err := run.AddDrawingAnchored(imgRef)
		if err != nil {
			log.Fatalln(err)
		}
		anchored.SetSize(measurement.Inch, measurement.Inch)
	}
}

// base64字符串转图片并保存
func base2img(base64Str string) (imgPath string) {
	imgPath = fmt.Sprintf("./html2word/image/%s", uuid.UUID()+".jpg")
	ddd, _ := base64.RawStdEncoding.DecodeString(base64Str)
	err := ioutil.WriteFile(imgPath, ddd, 0666)
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
