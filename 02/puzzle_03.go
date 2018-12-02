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

	twoTimes := 0
	threeTimes := 0

	scn := bufio.NewScanner(fh)

	for scn.Scan() {
		twoTimesHit, threeTimesHit := countTimes(countChars(scn.Text()))
		twoTimes += twoTimesHit
		threeTimes += threeTimesHit
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(twoTimes * threeTimes)
}

func countChars(str string) map[string]int {
	set := make(map[string]int)

	chars := strings.Split(str, "")
	for _, sym := range chars {
		if _, ok := set[sym]; ok {
			set[sym] += 1
		} else {
			set[sym] = 1
		}
	}

	return set
}

func countTimes(set map[string]int) (int, int) {
	twoTimes := 0
	threeTimes := 0

	for _, v := range set {
		switch v {
		case 2:
			twoTimes = 1
		case 3:
			threeTimes = 1
		}
	}

	return twoTimes, threeTimes
}
