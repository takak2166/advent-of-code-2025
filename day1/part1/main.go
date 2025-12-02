package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/takak2166/advent-of-code-2025/common"
)

const (
	initialPosition uint8 = 50
	maxDial         uint8 = 100
)

func main() {
	input, err := common.FetchInput(2025, 1)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	lines := strings.Split(input, "\n")
	var count uint = 0
	position := initialPosition
	for _, line := range lines {
		if line != "" {
			position = turn(position, line)
			// fmt.Println("rotation:", line, "position:", position, "count:", count)
			if position == 0 {
				count++
			}
		}
	}
	fmt.Println("Number of times dialed to 0:", count)
}

func turn(position uint8, rotation string) (nextPosition uint8) {
	if strings.HasPrefix(rotation, "R") {
		distance, _ := strconv.Atoi(strings.TrimPrefix(rotation, "R"))
		tmpPosition := int(position) + distance
		nextPosition = uint8(tmpPosition % int(maxDial))
		return
	}
	if strings.HasPrefix(rotation, "L") {
		distance, _ := strconv.Atoi(strings.TrimPrefix(rotation, "L"))
		tmpPosition := int(position) - distance
		for tmpPosition < 0 {
			tmpPosition += int(maxDial)
		}
		nextPosition = uint8(tmpPosition % int(maxDial))
		return
	}
	return 255
}
