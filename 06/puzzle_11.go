package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"math"
)

func main() {
	var coords [][]float64

	fh, err := os.Open("sample.txt")
	defer fh.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	scn := bufio.NewScanner(fh)

	for scn.Scan() {
		coords = append(coords, getCoords(scn.Text()))
	}

	check(scn.Err())

	fmt.Println(coords)
	// for i, c := range coords {
	// }


	// fmt.Printf("%f", manhattanDistance(coords[i], coords[j]))
}

func getCoords(str string) []float64 {
	var c []float64
	pts := strings.Split(str, ", ")

	a, err := strconv.ParseFloat(pts[0], 64)
	check(err)
	c = append(c, a)

	b, err := strconv.ParseFloat(pts[1], 64)
	check(err)
	c = append(c, b)

	return c
}

func manhattanDistance(a []float64, b []float64) float64 {
	s := float64(0)
	for i := 0; i < len(a); i += 1 {
		s += math.Abs(b[i] - a[i])
	}

	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
