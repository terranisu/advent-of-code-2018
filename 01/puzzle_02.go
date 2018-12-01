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

	// Create a set
	integerSet := make(map[int64]struct{})
	frequency := int64(0)

  // Run an infinite outer loop until the first frequency reached twice is found
  Outer:
		for {
			// Rewind the file pointer to the first line
			fileHandler.Seek(0, 0)
			scanner := bufio.NewScanner(fileHandler)

			// Iterate over file lines
			for scanner.Scan() {
				deviation, err := strconv.ParseInt(scanner.Text(), 10, 64)

				if err != nil {
					fmt.Fprintln(os.Stderr, "parsing string:", err)
				}

				frequency += deviation

				if _, ok := integerSet[frequency]; ok {
					break Outer
				}

				integerSet[frequency] = struct{}{}
			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}

	fmt.Println(frequency)
}
