package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	lineCountFlag := flag.Bool("l", false, "Count lines")
	wordCountFlag := flag.Bool("w", false, "Count words")
	byteCountFlag := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	filePath := flag.CommandLine.Arg(0)

	// #region Pipe/File handling
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File

	//The bitwise AND operation between stat.Mode() and os.ModeCharDevice checks if the ModeCharDevice bit is set in the FileMode.
	//If the result is non-zero, it means the ModeCharDevice bit is set, indicating that the file is a character device (terminal).
	//If the result is zero, it means the ModeCharDevice bit is not set, indicating that the file is not a character device (not a terminal).
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
	// #endregion

	if *lineCountFlag {
		fmt.Print(getLineCount(file))
	} else if *wordCountFlag {
		fmt.Print(getWordCount(file))
	} else if *byteCountFlag {
		fmt.Print(getByteCount(file))
	} else {
		fmt.Printf("%d  %d  %d", getLineCount(file), getWordCount(file), getByteCount(file))
	}
	fmt.Print(" ", filePath)

}

func getLineCount(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return -1
	}
	return lineCount
}
func getWordCount(file *os.File) int {
	scanner := bufio.NewScanner(file)
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		wordCount += len(strings.Fields(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return -1
	}
	return wordCount
}
func getByteCount(file *os.File) int {
	scanner := bufio.NewScanner(file)
	byteCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		byteCount += len(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return -1
	}
	return byteCount
}
