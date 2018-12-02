package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	boxes := strings.Split(string(data), "\n")

	for _, boxone := range boxes {
		for _, boxtwo := range boxes {
			if boxone == boxtwo {
				continue
			}

			correctBox, commonCharString := compareBoxSlice(boxone, boxtwo)
			if correctBox {
				fmt.Printf("Common characters from %v and %v are %v\n", boxone, boxtwo, commonCharString)
				os.Exit(0)
			}
		}
	}
}

func compareBoxSlice(boxone string, boxtwo string) (correctBox bool, commonCharString string) {

	boxoneSlice := strings.Split(boxone, "")
	boxtwoSlice := strings.Split(boxtwo, "")

	notsimilarCharCounter := 0

	for i := range boxoneSlice {
		if boxoneSlice[i] != boxtwoSlice[i] {
			notsimilarCharCounter++
			boxoneSlice[i] = ""
		}
	}

	if notsimilarCharCounter == 1 {
		fmt.Println("FOUND")
		correctBox = true
		commonCharString = strings.Join(boxoneSlice, "")
	}

	return correctBox, commonCharString

}
