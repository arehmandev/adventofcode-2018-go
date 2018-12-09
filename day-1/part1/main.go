package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	numbers := strings.Split(string(data), "\n")

	total := 0

	for _, number := range numbers {
		actual, _ := strconv.Atoi((string(number)))
		total += int(actual)
	}

	fmt.Println(total)
}
