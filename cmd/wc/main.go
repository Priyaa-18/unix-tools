package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if filename was provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: wc <filename>")
	}

	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Count lines
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print result
	fmt.Printf("%d %s\n", lineCount, filename)
}
