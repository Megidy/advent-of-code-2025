package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var res int

	for _, line := range lines {

		var (
			maxRight, maxLeft int
			// previousMaxRight  int
			previousMaxLeft int
			lineRune        = []rune(line)
		)

		for i := range lineRune {
			currRightStr := string(lineRune[i])
			currRight, _ := strconv.Atoi(currRightStr)

			if currRight > maxLeft {
				previousMaxLeft = maxLeft
				maxLeft = currRight
				// previousMaxRight = currRight
				maxRight = 0

				continue
			}

			if currRight > maxRight {

				maxRight = currRight
			}

		}
		if maxRight == 0 {
			maxRight = maxLeft
			maxLeft = previousMaxLeft
		}
		r, _ := strconv.Atoi(fmt.Sprintf("%d%d", maxLeft, maxRight))
		res += r
	}

	log.Println(res)
}
