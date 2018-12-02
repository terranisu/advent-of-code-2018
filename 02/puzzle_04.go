package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fh, err := os.Open("input.txt")
	defer fh.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	scn := bufio.NewScanner(fh)
	var list [250][]string
	i := 0

	for scn.Scan() {
		list[i] = strings.Split(scn.Text(), "")
		i += 1
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	j := 0
	Outer:
		for j < cap(list) {
			k := j
			for k < cap(list) {
				count, symbols := countDiffSymbols(list[j], list[k])

				if count == 1 {
					fmt.Println(symbols)
					break Outer
				}
				k += 1
			}
			j += 1
		}
}

func countDiffSymbols(a []string, b []string) (int, []string) {
	diffCount := 0
	symbols := make([]string, 26)

	for i, v := range a {
		if v != b[i] {
			diffCount += 1
		} else {
			symbols[i] = v
		}
	}

	return diffCount, symbols
}
