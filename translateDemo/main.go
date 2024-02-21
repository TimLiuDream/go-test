package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func main() {
	text := "Hello，World！"

	ja, err := translate(text, language.English.String(), language.Japanese.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("en: %s | ja: %s \n", text, ja)

	ch, err := translate(text, language.English.String(), language.Chinese.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("en: %s | ch: %s \n", text, ch)

	kr, err := translate(text, language.English.String(), language.Korean.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("en: %s | kr: %s \n", text, kr)
}

func translate(text string, src, dst string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: src,
			To:   dst,
		},
	)
	return translated, err
}
