package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"
)

const filePath = "/Users/tim/go/src/github.com/timliudream/golangtools/parsexml/小于号.opml"

type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	Title string `xml:"title"`
}

type Body struct {
	BodyOutlines []Outline `xml:"outline"`
}

type Outline struct {
	Text   string    `xml:"text,attr"`
	Childs []Outline `xml:"outline"`
}

func main() {
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	content := string(data)

	r, err := regexp.Compile("\".*(<+).*\"")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	content = r.ReplaceAllStringFunc(content, func(s string) string {
		return strings.Replace(s, "<", "", -1)
	})

	v := OPML{}

	sr := strings.NewReader(content)

	decoder := xml.NewDecoder(sr)
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity
	err = decoder.Decode(&v)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// decoder.Entity = xml.HTMLEntity
	// var strName string
	// for {
	// 	token, err := decoder.Token()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		break
	// 	}
	// 	switch t := token.(type) {
	// 	case xml.StartElement:
	// 		stelm := xml.StartElement(t)
	// 		fmt.Println("start: ", stelm.Name.Local)
	// 		strName = stelm.Name.Local
	// 	case xml.EndElement:
	// 		endelm := xml.EndElement(t)
	// 		fmt.Println("end: ", endelm.Name.Local)
	// 	case xml.CharData:
	// 		data := xml.CharData(t)
	// 		str := string(data)
	// 		switch strName {
	// 		case "text":
	// 			fmt.Println("text:", str)
	// 		}
	// 	}
	// }

	// err = xml.Unmarshal(data, &v)
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// 	return
	// }

	// // 构建树的所有路径，总路径数就是用例数
	// nodes := make([]string, 0) // 每个路径所经过的节点
	// paths := builtTreePaths(v.Body.BodyOutlines[0], nodes)
	// for _, path := range paths {
	// 	fmt.Println(path)
	// }
}

// 计算模块层级
func calculateXMLTreeMaxDepth(outline Outline) float64 {
	var depth float64
	for _, child := range outline.Childs {
		depth = math.Max(depth, calculateXMLTreeMaxDepth(child))
	}
	return depth + 1
}

// 求出树的所有路径
func builtTreePaths(outline Outline, nodes []string) (paths []string) {
	nodes = append(nodes, outline.Text)
	for _, child := range outline.Childs {
		if len(child.Childs) == 0 {
			nodes = append(nodes, child.Text)
			if paths == nil {
				paths = make([]string, 0)
			}
			paths = append(paths, buildNodesPath(nodes))
			nodes = nodes[:len(nodes)-1]
		} else {
			paths = append(paths, builtTreePaths(child, nodes)...)
		}
	}
	return
}

func buildNodesPath(nodes []string) (path string) {
	sb := strings.Builder{}
	for index, node := range nodes {
		sb.WriteString(node)
		if index != len(nodes)-1 {
			sb.WriteString("-")
		}
	}
	path = sb.String()
	return
}
