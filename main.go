package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPath := flag.String("in", "", "Path to the input SRT file")
	outputPath := flag.String("out", "", "Path to the output SRT file")
	fromLang := flag.String("from", "", "Source language code (e.g. en)")
	toLang := flag.String("to", "", "Target language code (e.g. se)")
	flag.Parse()

	if *inputPath == "" || *outputPath == "" || *fromLang == "" || *toLang == "" {
		fmt.Println(
			"Usage: srtranslate -in input.srt -out output.srt -from en -to se")
		os.Exit(1)
	}

	inputFile, err := os.Open(*inputPath)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	subtitles, err := parseSRT(inputFile)
	if err != nil {
		fmt.Printf("Error parsing SRT file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting translation...")
	translatedSubtitles, err := translateSubtitles(subtitles, *fromLang, *toLang)
	if err != nil {
		fmt.Printf("Error during translation: %v\n", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	err = writeSRT(outputFile, translatedSubtitles)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Translation completed successfully!")
}
