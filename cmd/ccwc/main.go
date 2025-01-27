package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type stats struct {
	lineCount int64
	wordCount int64
	byteCount int64
	charCount int64
}

func GetStats(file *os.File) stats {
	reader := bufio.NewReader(file)
	stats := stats{}
	endOfWord := false
	for {
		char, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if endOfWord {
					stats.wordCount++
				}
				break
			}
			log.Fatal(err)
		}
		stats.byteCount += int64(size)
		stats.charCount++
		if unicode.IsSpace(char) {
			if endOfWord {
				stats.wordCount++
				endOfWord = false
			}
		} else {
			endOfWord = true
		}
		if char == '\n' {
			stats.lineCount++
		}
	}
	return stats
}

func main() {

	lineCountFlag := flag.Bool("l", false, "Count lines")
	wordCountFlag := flag.Bool("w", false, "Count words")
	byteCountFlag := flag.Bool("c", false, "Count bytes")
	charCountFlag := flag.Bool("m", false, "Count characters")
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

	stats := GetStats(file)

	if *lineCountFlag {
		fmt.Print(stats.lineCount)
	} else if *wordCountFlag {
		fmt.Print(stats.wordCount)
	} else if *byteCountFlag {
		fmt.Print(stats.byteCount)
	} else if *charCountFlag {
		fmt.Print(stats.charCount)
	} else {
		fmt.Printf("%d  %d  %d", stats.lineCount, stats.wordCount, stats.byteCount)
	}
	fmt.Printf(" %s \n", filePath)

}

func GetStatsAlternate(file *os.File) stats {
	scanner := bufio.NewScanner(file)
	stats := stats{}
	for scanner.Scan() {
		line := scanner.Text()
		stats.lineCount++
		stats.wordCount += int64(len(strings.Fields(line)))
		stats.byteCount += int64(len(line))
		stats.charCount += int64(len([]rune(line)))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stats
}
