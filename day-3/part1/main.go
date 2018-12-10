package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	grids := strings.Split(string(data), "\n")

	coordinateMap := make(map[coordinate]int)

	for _, grid := range grids {
		gridlines := strings.Split(string(grid), " ")

		leftspace, _ := strconv.Atoi((strings.Split(string(gridlines[2]), ",")[0]))
		topspace, _ := strconv.Atoi((strings.Replace(strings.Split(string(gridlines[2]), ",")[1], ":", "", -1)))
		gridwidth, _ := strconv.Atoi((strings.Split(string(gridlines[3]), "x")[0]))
		gridlength, _ := strconv.Atoi((strings.Split(string(gridlines[3]), "x")[1]))

		for indexw := 0; indexw < gridwidth; indexw++ {

			for indexl := 0; indexl < gridlength; indexl++ {

				newcoord := coordinate{x: (leftspace + indexw), y: (topspace + indexl)}
				coordinateMap[newcoord]++
			}

		}

	}

	fabricCount := 0
	for _, value := range coordinateMap {
		if value > 1 {
			fabricCount++
		}
	}
	fmt.Println(fabricCount)
}
