package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/takak2166/advent-of-code-2025/common"
)

func main() {
	input, err := common.FetchInput(2025, 2)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\n", "")

	rangeStrs := strings.Split(input, ",")
	ranges := make([][]uint, len(rangeStrs))
	for i, rangeStr := range rangeStrs {
		if rangeStr != "" {
			numStrs := strings.Split(rangeStr, "-")
			min, _ := strconv.Atoi(numStrs[0])
			max, _ := strconv.Atoi(numStrs[1])
			ranges[i] = []uint{uint(min), uint(max)}
		}
	}

	var ans uint = 0
	for _, r := range ranges {
		for j := r[0]; j <= r[1]; j++ {
			for n := range 6 {
				if n > 0 && uint(math.Pow10(2*n)) <= j && j < uint(math.Pow10(2*n+1)) && j%uint((math.Pow10(2*n+1)-1)/9) == 0 {
					ans += j
					fmt.Println("pattern1-1, found: ", j, "ans: ", ans)
				} else if uint(math.Pow10(2*n+1)) <= j && j < uint(math.Pow10(2*n+2)) && j%uint((math.Pow10(2*n+2)-1)/9) == 0 {
					ans += j
					fmt.Println("pattern1-2, found: ", j, "ans: ", ans)
				} else if uint(math.Pow10(2*n+1)) <= j && j < uint(math.Pow10(2*n+2)) && j%uint((math.Pow10(n+1))+1) == 0 {
					ans += j
					fmt.Println("pattern2, found: ", j, "ans: ", ans)
				} else if n > 0 && uint(math.Pow10(2*n+1)) <= j && j < uint(math.Pow10(2*n+2)) && j%uint((math.Pow10(2*n+2)-1)/99) == 0 {
					ans += j
					fmt.Println("pattern3, found: ", j, "ans: ", ans)
				} else if n == 4 && uint(math.Pow10(2*n)) <= j && j < uint(math.Pow10(2*n+1)) && j%1001001 == 0 {
					ans += j
					fmt.Println("pattern4, found: ", j, "ans: ", ans)
				}
			}
		}
	}

	fmt.Println("answer: ", ans)
}
