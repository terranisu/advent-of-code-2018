package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func main() {
	lines, err := readLines("input.txt")
	check(err)

	acc := reactive(lines[0])
	var str string
	var noCharAcc []string

	i := 1
	for i <= 26 {
		ch := toChar(i)
		var noChar []string

		for _, v := range acc {
			if v == ch || v == strings.ToUpper(ch) {
				continue
			}
			noChar = append(noChar, v)
		}

		i += 1
		str = strings.Join(noChar, "")
		noCharAcc = reactive(str)
		fmt.Println(ch, len(noCharAcc))
	}
}
// 4944
func toChar(i int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	return alphabet[i-1:i]
}

func reactive(bytes string) []string {
	var acc []string
	var curr string
	var next string

	i := 0
	for i < len(bytes) - 1 {
		if len(acc) == 0 {
			acc = append(acc, string(bytes[0]))
		}

		curr = acc[len(acc)-1]
		next = string(bytes[i + 1])

		// fmt.Println(acc, curr, next)

		if (curr == strings.ToLower(curr) && strings.ToUpper(curr) == next) ||
			(curr == strings.ToUpper(curr) && strings.ToLower(curr) == next) {
			acc = acc[:len(acc)-1]
			// fmt.Println("DESTROY")
		} else {
			acc = append(acc, next)
		}
		i += 1
	}

	return acc
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  check(err)
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
