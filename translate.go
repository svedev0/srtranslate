package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func translateSubtitles(subtitles []Subtitle, fromLang, toLang string) ([]Subtitle, error) {
	translatedSubtitles := make([]Subtitle, len(subtitles))
	lastSubIdx := subtitles[len(subtitles)-1].Index

	supportedLanguages := getSupportedLanguages()
	for _, lang := range supportedLanguages {
		if lang == fromLang {
			return nil, fmt.Errorf("unsupported source language code '%s'", fromLang)
		}
		if lang == toLang {
			return nil, fmt.Errorf("unsupported target language code '%s'", toLang)
		}
	}

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

func getSupportedLanguages() []string {
	return []string{
		"af",    // Afrikaans
		"sq",    // Albanian
		"am",    // Amharic
		"ar",    // Arabic
		"hy",    // Armenian
		"as",    // Assamese
		"ay",    // Aymara
		"az",    // Azerbaijani
		"bm",    // Bambara
		"eu",    // Basque
		"be",    // Belarusian
		"bn",    // Bengali
		"bs",    // Bosnian
		"bg",    // Bulgarian
		"ca",    // Catalan
		"zh-CN", // Chinese (Simplified)
		"zh-TW", // Chinese (Traditional)
		"co",    // Corsican
		"hr",    // Croatian
		"cs",    // Czech
		"da",    // Danish
		"dv",    // Dhivehi
		"nl",    // Dutch
		"en",    // English
		"eo",    // Esperanto
		"et",    // Estonian
		"ee",    // Ewe
		"fi",    // Finnish
		"fr",    // French
		"fy",    // Frisian
		"gl",    // Galician
		"ka",    // Georgian
		"de",    // German
		"el",    // Greek
		"gn",    // Guarani
		"gu",    // Gujarati
		"ha",    // Hausa
		"he",    // Hebrew
		"hi",    // Hindi
		"hu",    // Hungarian
		"is",    // Icelandic
		"ig",    // Igbo
		"id",    // Indonesian
		"ga",    // Irish
		"it",    // Italian
		"ja",    // Japanese
		"kn",    // Kannada
		"kk",    // Kazakh
		"km",    // Khmer
		"rw",    // Kinyarwanda
		"ko",    // Korean
		"ku",    // Kurdish
		"ky",    // Kyrgyz
		"lo",    // Lao
		"la",    // Latin
		"lv",    // Latvian
		"ln",    // Lingala
		"lt",    // Lithuanian
		"lg",    // Luganda
		"lb",    // Luxembourgish
		"mk",    // Macedonian
		"mg",    // Malagasy
		"ms",    // Malay
		"ml",    // Malayalam
		"mt",    // Maltese
		"mi",    // Maori
		"mr",    // Marathi
		"mn",    // Mongolian
		"my",    // Myanmar (Burmese)
		"ne",    // Nepali
		"no",    // Norwegian
		"om",    // Oromo
		"ps",    // Pashto
		"fa",    // Persian
		"pl",    // Polish
		"pt",    // Portuguese
		"pa",    // Punjabi
		"qu",    // Quechua
		"ro",    // Romanian
		"ru",    // Russian
		"sm",    // Samoan
		"sa",    // Sanskrit
		"gd",    // Scots Gaelic
		"sr",    // Serbian
		"st",    // Sesotho
		"sn",    // Shona
		"sd",    // Sindhi
		"sk",    // Slovak
		"sl",    // Slovenian
		"so",    // Somali
		"es",    // Spanish
		"su",    // Sundanese
		"sw",    // Swahili
		"sv",    // Swedish
		"tl",    // Tagalog (Filipino)
		"tg",    // Tajik
		"ta",    // Tamil
		"tt",    // Tatar
		"te",    // Telugu
		"th",    // Thai
		"ti",    // Tigrinya
		"ts",    // Tsonga
		"tr",    // Turkish
		"tk",    // Turkmen
		"uk",    // Ukrainian
		"ur",    // Urdu
		"ug",    // Uyghur
		"uz",    // Uzbek
		"vi",    // Vietnamese
		"cy",    // Welsh
		"xh",    // Xhosa
		"yi",    // Yiddish
		"yo",    // Yoruba
		"zu",    // Zulu
	}
}
