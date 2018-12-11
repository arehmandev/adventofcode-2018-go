package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	input := string(data)

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

	fmt.Println(len(input))

}
