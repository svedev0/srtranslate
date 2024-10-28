package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func translateSubtitles(subtitles []Subtitle, fromLang, toLang string) ([]Subtitle, error) {
	translatedSubtitles := make([]Subtitle, len(subtitles))
	lastSubIdx := subtitles[len(subtitles)-1].Index

	_, err := language.Parse(fromLang)
	if err != nil {
		return nil, fmt.Errorf("invalid source language code '%s': %v", fromLang, err)
	}
	_, err = language.Parse(toLang)
	if err != nil {
		return nil, fmt.Errorf("invalid target language code '%s': %v", toLang, err)
	}

	for i, sub := range subtitles {
		translated, err := gtranslate.TranslateWithParams(
			sub.Text, gtranslate.TranslationParams{From: fromLang, To: toLang})
		if err != nil {
			return nil, fmt.Errorf("error translating subtitle %d: %v", sub.Index, err)
		}

		translatedSubtitles[i] = Subtitle{
			Index:     sub.Index,
			Timestamp: sub.Timestamp,
			Text:      translated,
		}

		percentStr := fmt.Sprintf("%.2f", 100*(float64(i)/float64(lastSubIdx)))
		fmt.Printf("Progress %s%%\n", percentStr)
	}

	return translatedSubtitles, nil
}
