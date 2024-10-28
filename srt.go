package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Subtitle struct {
	Index     int
	Timestamp string
	Text      string
}

func parseSRT(r io.Reader) ([]Subtitle, error) {
	var subtitles []Subtitle
	var current Subtitle

	scanner := bufio.NewScanner(r)
	reTimestamp := regexp.MustCompile(`^\d{2}:\d{2}:\d{2},\d{3} --> \d{2}:\d{2}:\d{2},\d{3}$`)

	// 0: expecting index
	// 1: expecting timestamp
	// 2: expecting text
	state := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		switch state {
		case 0:
			if line == "" {
				continue
			}

			var idx int
			_, err := fmt.Sscanf(line, "%d", &idx)
			if err != nil {
				return nil, fmt.Errorf("expected subtitle index, got '%s'", line)
			}

			current = Subtitle{Index: idx}
			state = 1

		case 1:
			if !reTimestamp.MatchString(line) {
				return nil, fmt.Errorf("expected timestamp, got '%s'", line)
			}
			current.Timestamp = line
			state = 2

		case 2:
			if line == "" {
				subtitles = append(subtitles, current)
				state = 0
			} else {
				if current.Text == "" {
					current.Text = line
				} else {
					current.Text += "\n" + line
				}
			}
		}
	}

	if state == 2 && current.Text != "" {
		subtitles = append(subtitles, current)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subtitles, nil
}

func writeSRT(w io.Writer, subtitles []Subtitle) error {
	for _, sub := range subtitles {
		_, err := fmt.Fprintf(w, "%d\n%s\n%s\n\n", sub.Index, sub.Timestamp, sub.Text)
		if err != nil {
			return err
		}
	}

	return nil
}
