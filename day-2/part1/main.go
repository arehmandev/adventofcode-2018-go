package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	boxes := strings.Split(string(data), "\n")

	twocount := 0
	threecount := 0

	for _, box := range boxes {
		letterDict := make(map[string]int)
		for _, letter := range box {
			letterDict[string(letter)]++
		}

		twocountFound := false
		threecountFound := false

		for _, count := range letterDict {
			if twocountFound && threecountFound {
				break
			}

			if count == 2 && !twocountFound {
				twocount++
				twocountFound = true
			}
			if count == 3 && !threecountFound {
				threecount++
				threecountFound = true
			}

		}

	}

	fmt.Println(twocount * threecount)
}
