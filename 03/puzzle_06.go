package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type Rectangle struct {
	id string
	x, y, h, w int64
}

func main() {
	fh, err := os.Open("input.txt")
	defer fh.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	scn := bufio.NewScanner(fh)
	rects := []Rectangle{}
	idsSet := make(map[string]struct{})

	for scn.Scan() {
		parts := strings.Split(scn.Text(), " ")

		id := parts[0][1:]
		osX, osY := parseOffset(parts[2])
		sX, sY := parseSize(parts[3])

		rects = append(rects, Rectangle{id, osX, osY, sY, sX})
		idsSet[id] = struct{}{}
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	idSet := make(map[string]struct{})

	for i := 0; i < len(rects); i += 1 {
		for j := 0; j < len(rects); j += 1 {
			if i == j {
				continue
			} else if isOverlapped(rects[i], rects[j]) {
				idSet[rects[i].id] = struct{}{}
				idSet[rects[j].id] = struct{}{}
				break
			}
		}
	}

	for k, _ := range idsSet {
		if _, ok := idSet[k]; !ok {
			fmt.Println(k)
		}
	}
}

func isOverlapped(r1 Rectangle, r2 Rectangle) bool {
	if r1.x + r1.w < r2.x || r2.x + r2.w < r1.x || r1.y + r1.h < r2.y || r2.y + r2.h < r1.y {
		return false
	}

	return true
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
