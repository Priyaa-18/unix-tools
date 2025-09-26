package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Define flags
	countLines := flag.Bool("l", false, "count lines")
	countWords := flag.Bool("w", false, "count words")
	countBytes := flag.Bool("c", false, "count bytes")
	countChars := flag.Bool("m", false, "count characters")
	flag.Parse()

	// Check if filename was provided
	if flag.NArg() < 1 {
		log.Fatal("Usage: wc [options] <filename>")
	}
	filename := flag.Arg(0)

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Counters
	lines, words, bytes, chars := 0, 0, 0, 0

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		bytes += len(line) + 1 // +1 for newline
		chars += utf8.RuneCountInString(line) + 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// If no flags given, mimic default wc (print all)
	if !*countLines && !*countWords && !*countBytes && !*countChars {
		fmt.Printf("%d %d %d %d %s\n", lines, words, bytes, chars, filename)
		return
	}

	// Print only requested flags
	if *countLines {
		fmt.Printf("%d ", lines)
	}
	if *countWords {
		fmt.Printf("%d ", words)
	}
	if *countBytes {
		fmt.Printf("%d ", bytes)
	}
	if *countChars {
		fmt.Printf("%d ", chars)
	}
	fmt.Printf("%s\n", filename)
}
