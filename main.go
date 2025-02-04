package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Improved regex to capture full path from first "/" to extension
	re := regexp.MustCompile(`/[^ \t\n]+?\.[a-zA-Z0-9]{3}`)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // Read line by line
		if err != nil && len(line) == 0 {
			break // Exit loop if no more content
		}

		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
