package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

var jsonStr = `{
	"name": "Tim",
	"age": 31,
	"gender": "male"
}`

func main() {
	normalDecode()
	fmt.Println("------------------")
	embeddedStructDecode()
	fmt.Println("------------------")
	decodeErrorHandle()
	fmt.Println("------------------")
	metadataDecode()
	fmt.Println("------------------")
	omitemptyDecode()
	fmt.Println("------------------")
	remainDataDecode()
	fmt.Println("------------------")
	tagDecode()
	fmt.Println("------------------")
	weaklyTypedInputDecode()
	fmt.Println("------------------")
	jsonDecode()
}

func normalDecode() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	// 翻译：此输入可以来自任何地方，但通常来自诸如解码 JSON 之类的东西，我们最初不太确定结构。
	input := map[string]interface{}{
		"name":   "Tim",
		"age":    31,
		"emails": []string{"one@gmail.com", "two@gmail.com", "three@gmail.com"},
		"extra": map[string]string{
			"twitter": "Tim",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}

func embeddedStructDecode() {
	// Squashing multiple embedded structs is allowed using the squash tag.
	// This is demonstrated by creating a composite struct of multiple types
	// and decoding into it. In this case, a person can carry with it both
	// a Family and a Location, as well as their own FirstName.
	// 翻译：使用 squash 标签允许压缩多个嵌入式结构。通过创建多种类型的复合结构并对其进行解码来演示此功能。在这种情况下，人可以携带一个家庭和一个位置，以及他们自己的名字。
	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		Family    `mapstructure:",squash"`
		Location  `mapstructure:",squash"`
		FirstName string
	}

	input := map[string]interface{}{
		"FirstName": "Tim",
		"LastName":  "Liu",
		"City":      "China, Guangdong",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s, %s\n", result.FirstName, result.LastName, result.City)
}

func decodeErrorHandle() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	// 翻译：此输入可以来自任何地方，但通常来自诸如解码 JSON 之类的东西，我们最初不太确定结构。
	input := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func metadataDecode() {
	type Person struct {
		Name string
		Age  int
	}

	// This input can come from anywhere, but typically comes from something like decoding JSON where we're not quite sure of the struct initially.
	// 翻译：此输入可以来自任何地方，但通常来自诸如解码 JSON 之类的东西，我们最初不太确定结构。
	input := map[string]interface{}{
		"name":  "Tim",
		"age":   31,
		"email": "one@gmail.com",
	}

	// For metadata, we make a more advanced DecoderConfig so we can more finely configure the decoder that is used. In this case, we just tell the decoder we want to track metadata.
	// 翻译：对于元数据，我们制作了一个更高级的 DecoderConfig，以便我们可以更细致地配置所使用的解码器。在这种情况下，我们只是告诉解码器我们想要跟踪元数据。
	var md mapstructure.Metadata
	var result Person
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err = decoder.Decode(input); err != nil {
		panic(err)
	}

	fmt.Printf("value: %#v, Unused keys: %#v\n", result, md.Unused)
}

func omitemptyDecode() {
	// Add omitempty annotation to avoid map keys for empty values
	// 翻译：添加 omitempty 注释以避免空值的映射键
	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		*Family   `mapstructure:",omitempty"`
		*Location `mapstructure:",omitempty"`
		Age       int
		FirstName string
	}

	result := &map[string]interface{}{}
	input := Person{FirstName: "Somebody"}
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", result)
}

func remainDataDecode() {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	// 翻译：请注意，结构类型中定义的 mapstructure 标签可以指示将值映射到哪些字段。
	type Person struct {
		Name  string
		Age   int
		Other map[string]interface{} `mapstructure:",remain"`
	}

	input := map[string]interface{}{
		"name":   "Tim",
		"age":    31,
		"email":  "one@gmail.com",
		"gender": "male",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}

func tagDecode() {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	// 翻译：请注意，结构类型中定义的 mapstructure 标签可以指示将值映射到哪些字段。
	type Person struct {
		Name string `mapstructure:"person_name"`
		Age  int    `mapstructure:"person_age"`
	}

	input := map[string]interface{}{
		"person_name": "Tim",
		"person_age":  31,
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}

func weaklyTypedInputDecode() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON, generated by a weakly typed language such as PHP.
	// 翻译：此输入可以来自任何地方，但通常来自诸如解码 JSON 之类的东西，由 PHP 等弱类型语言生成。
	input := map[string]interface{}{
		"name": 123,  // number => string
		"age":  "31", // string => number
		"emails": map[string]interface{}{
			//"one@gmail.com": true,//* 'Emails[0]' expected type 'string', got unconvertible type 'map[string]interface {}', value: 'map[one@gmail.com:true]'
			//"emails": []string{
			//	"one@gmail.com",
			//},* 'Emails[0]' expected type 'string', got unconvertible type 'map[string]interface {}', value: 'map[emails:[one@gmail.com]]'
		}, // empty map => empty array
	}

	var result Person
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}

func jsonDecode() {
	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		panic(err)
	}

	var result Person
	err = mapstructure.Decode(m, &result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", result)
}
