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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getFileCounts(file *os.File) (bytesCount int, linesCount int, wordsCount int, charsCount int) {
	bytesCount = 0
	linesCount = 0
	wordsCount = 0
	charsCount = 0

	inWord := false

	reader := bufio.NewReader(file)

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
	flag.Parse()
	fileName := flag.Args()[0]
	if !getBytesCount && !getLinesCount && !getWordsCount && !getCharsCount {
		getBytesCount = true
		getLinesCount = true
		getWordsCount = true
		getCharsCount = true
	}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	bytesCount, linesCount, wordsCount, charsCount := getFileCounts(file)

	var output []string

	if getBytesCount {
		output = append(output, strconv.Itoa(bytesCount))
	}
	if getLinesCount {
		output = append(output, strconv.Itoa(linesCount))
	}
	if getWordsCount {
		output = append(output, strconv.Itoa(wordsCount))
	}
	if getCharsCount {
		output = append(output, strconv.Itoa(charsCount))
	}
	output = append(output, filepath.Base(fileName))

	fmt.Println(strings.Join(output, " "))
}
