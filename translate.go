package main 

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"github.com/mind1949/googletrans"
    
)

// free ref: go get -u github.com/mind1949/googletrans
// org ref: https://cloud.google.com/translate/docs/basic/translating-text#translate_translate_text-go
func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
			return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
			return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
			return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
			return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

func detect_language(filename string) {

	detected, err := googletrans.Detect(filename)
	if err != nil {
		panic(err)
	}

	format := "language: %q, confidence: %0.2f\n"
	fmt.Printf(format, detected.Lang, detected.Confidence)
}

func translate_text(filename string) {
	params := googletrans.TranslateParams{
		Src:  "auto",
		Dest: language.SimplifiedChinese.String(),
		// Text: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. ",
		Text: filename,
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
}