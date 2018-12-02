package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var data = `-1
-7
-5
-16`

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	numbers := strings.Split(string(data), "\n")

	total := 0
	totaldict := make(map[int]int)

	for {
		breaker := false
		for _, number := range numbers {
			actual, _ := strconv.Atoi((string(number)))
			total += int(actual)

			if _, ok := totaldict[total]; ok {
				fmt.Println("Total reached twice:", total)
				breaker = true
				break
			} else {
				totaldict[total] = 1
			}
		}
		if breaker {
			break
		}
	}

	fmt.Println(total)
}
