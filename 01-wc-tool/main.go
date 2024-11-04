package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

var (
	getBytesCount bool
	getLinesCount bool
	getWordsCount bool
	getCharsCount bool
)

func getFileCounts(reader *bufio.Reader) (bytesCount int, linesCount int, wordsCount int, charsCount int) {
	bytesCount = 0
	linesCount = 0
	wordsCount = 0
	charsCount = 0

	inWord := false

	for {
		char, size, err := reader.ReadRune()
		if err != nil && err.Error() == "EOF" {
			if inWord {
				wordsCount++
			}
			break
		}
		if string(char) == "\n" {
			linesCount++
		}
		if unicode.IsSpace(char) {
			if inWord {
				wordsCount++
			}
			inWord = false
		} else {
			inWord = true
		}
		bytesCount += size
		charsCount++
	}

	return bytesCount, linesCount, wordsCount, charsCount
}

func main() {
	flag.BoolVar(&getBytesCount, "c", false, "print the byte counts")
	flag.BoolVar(&getLinesCount, "l", false, "print the newline counts")
	flag.BoolVar(&getWordsCount, "w", false, "print the word counts")
	flag.BoolVar(&getCharsCount, "m", false, "print the character counts")
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: ccwc [OPTION]... [FILE]...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !getBytesCount && !getLinesCount && !getWordsCount && !getCharsCount {
		getBytesCount = true
		getLinesCount = true
		getWordsCount = true
	}

	var reader *bufio.Reader
	var fileName string

	if len(flag.Args()) != 0 {
		fileName = flag.Args()[0]
		file, err := os.Open(fileName)
		if err != nil {
			os.Stderr.WriteString("Error: " + err.Error() + "\n")
			os.Exit(1)
		}
		defer file.Close()

		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	bytesCount, linesCount, wordsCount, charsCount := getFileCounts(reader)

	var output []string

	if getLinesCount {
		output = append(output, strconv.Itoa(linesCount))
	}
	if getWordsCount {
		output = append(output, strconv.Itoa(wordsCount))
	}
	if getBytesCount {
		output = append(output, strconv.Itoa(bytesCount))
	}
	if getCharsCount {
		output = append(output, strconv.Itoa(charsCount))
	}
	if fileName != "" {
		output = append(output, filepath.Base(fileName))
	}

	fmt.Println(strings.Join(output, " "))
}
