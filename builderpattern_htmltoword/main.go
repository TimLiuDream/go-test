package main

import "fmt"

type WordElement struct {
	Value string
}

func (e *WordElement) SetValue(value string) {
	e.Value = value
}

func (e *WordElement) GetValue() string {
	return e.Value
}

type Builder interface {
	SetValue(value string) Builder
	Build() *WordElement
}

type WordElementBuilder struct {
	e *WordElement
}

func (wordElementBuilder *WordElementBuilder) SetValue(value string) Builder {
	if wordElementBuilder.e == nil {
		wordElementBuilder.e = &WordElement{}
	}
	wordElementBuilder.e.SetValue(value)
	return wordElementBuilder
}

func (wordElementBuilder *WordElementBuilder) Build() *WordElement {
	return wordElementBuilder.e
}

type Director struct {
	builder Builder
}

func (d Director) Create(value string) *WordElement {
	return d.builder.SetValue(value).Build()
}

func main() {
	var builder Builder = &WordElementBuilder{}
	var diretor *Director = &Director{builder: builder}
	var h *WordElement = diretor.Create("测试导出1")
	fmt.Println(h.GetValue())
}
