package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

func main() {
	fh, err := os.Open("sorted_input.txt")
	defer fh.Close()

	check(err)

	scn := bufio.NewScanner(fh)
	sleptTotal := make(map[string]int)
	sleptByMins := make(map[string][]int)
	curId := ""
	awakeAt := int64(0)
	sleepAt := int64(0)

	for scn.Scan() {
		parts := strings.Split(scn.Text(), "] ")

		if strings.HasPrefix(parts[1], "Guard") {
			re := regexp.MustCompile("\\d+")
			id := re.FindString(parts[1])
			if _, ok := sleptTotal[id]; !ok {
				sleptTotal[id] = 0
				i := int64(0)
				for i < 59 {
					sleptByMins[id] = append(sleptByMins[id], 0)
					i += 1
				}
			}
			curId = id
			awakeAt = 0
			sleepAt = 0
		} else if strings.HasPrefix(parts[1], "wakes") {
			min, err := strconv.ParseInt(parts[0][len(parts[0])-2:], 10, 0)
			check(err)
			awakeAt = min
			if sleepAt > 0 {
				i := sleepAt
				for i < awakeAt {
					sleptByMins[curId][i] += 1
					sleptTotal[curId] += 1
					i += 1
				}
			}
		} else {
			min, err := strconv.ParseInt(parts[0][len(parts[0])-2:], 10, 0)
			check(err)
			sleepAt = min
		}
	}

	id := ""
	max := 0
	for k, v := range sleptTotal {
		if v > max {
			max = v
			id = k
		}
	}

	maxMins := 0
	minute := 0
	for k, v := range sleptByMins[id] {
		if v > maxMins {
			maxMins = v
			minute = k
		}
	}

	idAsInt, err := strconv.ParseInt(id, 10, 0)
	fmt.Println(int(idAsInt) * minute)

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
