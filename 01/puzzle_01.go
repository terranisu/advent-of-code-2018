package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	fileHandler, err := os.Open("input.txt")
	defer fileHandler.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	frequency := int64(0)
	scanner := bufio.NewScanner(fileHandler)

	for scanner.Scan() {
		deviation, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			fmt.Fprintln(os.Stderr, "parsing string:", err)
		}

		frequency += deviation
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(frequency)
}
