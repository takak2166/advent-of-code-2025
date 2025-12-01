package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	initialPosition uint8 = 50
	maxDial               = 100
)

func main() {
	session := os.Getenv("SESSION")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://adventofcode.com/2025/day/1/input", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	input := string(body)
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
	return
}

func turn(position uint8, rotation string) (nextPosition uint8) {
	if strings.HasPrefix(rotation, "R") {
		distance, _ := strconv.Atoi(strings.TrimPrefix(rotation, "R"))
		tmpPosition := int(position) + distance
		nextPosition = uint8(tmpPosition % maxDial)
		return
	}
	if strings.HasPrefix(rotation, "L") {
		distance, _ := strconv.Atoi(strings.TrimPrefix(rotation, "L"))
		tmpPosition := int(position) - distance
		for tmpPosition < 0 {
			tmpPosition += maxDial
		}
		nextPosition = uint8(tmpPosition % maxDial)
		return
	}
	return 255
}
