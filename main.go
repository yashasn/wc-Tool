package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	lineCountFlag := flag.Bool("l", false, "Count lines")
	wordCountFlag := flag.Bool("w", false, "Count words")
	byteCountFlag := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}
	filePath := args[0]

	// Open file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

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
		fmt.Println("Error reading file:", err)
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
		fmt.Println("Error reading file:", err)
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
		fmt.Println("Error reading file:", err)
		return -1
	}
	return byteCount
}
