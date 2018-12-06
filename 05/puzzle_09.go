package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	check(err)

	var acc []string
	var curr string
	var next string

	i := 0
	for i < len(bytes) - 3 {
		if len(acc) == 0 {
			acc = append(acc, string(bytes[0]))
		}

		curr = acc[len(acc)-1]
		next = string(bytes[i + 1])

		// fmt.Println(acc, curr, next)

		if (curr == strings.ToLower(curr) && strings.ToUpper(curr) == next) ||
			(curr == strings.ToUpper(curr) && strings.ToLower(curr) == next) {
			fmt.Println(curr, next)
			acc = acc[:len(acc)-1]
			j += 1
			// fmt.Println("DESTROY")
		} else {
			acc = append(acc, next)
		}
		i += 1
	}

	// fmt.Println(i)
	fmt.Println(len(acc))
	fmt.Println(j)
	// fmt.Println(string(acc))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
