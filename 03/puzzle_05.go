package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)


func main() {
	fh, err := os.Open("input.txt")
	defer fh.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	scn := bufio.NewScanner(fh)
	fbc := make(map[string]int)

	for scn.Scan() {
		parts := strings.Split(scn.Text(), " ")

		osX, osY := parseOffset(parts[2])
		sX, sY := parseSize(parts[3])

		for i := osX; i < osX + sX; i += 1 {
			for j := osY; j < osY + sY; j += 1 {
				key := fmt.Sprintf("x%dy%d", i, j)
				if _, ok := fbc[key]; ok {
					fbc[key] += 1
				} else {
					fbc[key] = 1
				}
			}
		}
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	total := 0
	for _, v := range fbc {
		if v > 1 {
			total += 1
		}
	}

	fmt.Println(total)

}

func parseOffset(offset string) (int64, int64) {
	parts := strings.Split(offset[:len(offset) - 1], ",")

	x, _ := strconv.ParseInt(parts[0], 10, 0)
	y, _ := strconv.ParseInt(parts[1], 10, 0)

	return x, y
}

func parseSize(size string) (int64, int64) {
	parts := strings.Split(size, "x")

	x, _ := strconv.ParseInt(parts[0], 10, 0)
	y, _ := strconv.ParseInt(parts[1], 10, 0)

	return x, y
}
