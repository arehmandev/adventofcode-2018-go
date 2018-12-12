package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	stringInput := string(data)

	lowestVal := len(stringInput)
	var loweststringletter string

	for j := 65; j < 65+26; j++ {

		stringletter := string(rune(j))
		capStringLetter := string(rune(j + 32))

		replacedInput := strings.Replace(stringInput, stringletter, "", -1)
		replacedInput = strings.Replace(replacedInput, capStringLetter, "", -1)

		result := reactPolymer(replacedInput)
		fmt.Println(stringletter, len(result))

		if len(result) < lowestVal {
			lowestVal = len(result)
			loweststringletter = stringletter
		}
	}

	fmt.Println("LowestVal - ", loweststringletter, lowestVal)
}

func reactPolymer(input string) string {
	found := true

	for found {
		found = false
		for i := 0; i < len(input)-1; i++ {
			if input[i] == input[i+1]+32 || input[i] == input[i+1]-32 {
				input = input[:i] + input[i+2:]
				found = true
			}
		}
	}

	return input
}
