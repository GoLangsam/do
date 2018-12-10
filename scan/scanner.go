// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func linesOfWords(out chan<- []string, r io.Reader, skipEmptyLines bool) {
	defer close(out)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		scanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		scanner.Split(bufio.ScanWords)

		words := []string{}
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Word-Scanner encountered:", err)
		}

		if skipEmptyLines && len(words) < 1 {
			continue
		}

		out <- words
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Line-Scanner encountered:", err)
	}
	return
}

// LinesOfWords returns a channel to receive the slice of words as found in each line.
func LinesOfWords(r io.Reader, skipEmptyLines bool) <-chan []string {
	cha := make(chan []string)
	go linesOfWords(cha, r, skipEmptyLines)
	return cha
}

// WordsPerLine returns a channel to receive the slice of words as found in each line.
// Only non-empty lines are considered, thus all no empty
func WordsPerLine(r io.Reader) <-chan []string {
	return LinesOfWords(r, true)
}
