package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	aAndB()
}

func aAndB() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var currPosition int = 50
	var zeroCounter int

	for _, line := range lines {
		line = strings.TrimSpace(line)

		direction := line[:1]
		positionStr := line[1:]
		position, _ := strconv.Atoi(positionStr)

		if position > 100 {
			position = position % 100
		}

		switch direction {
		case "L":
			currPosition -= position
			if currPosition < 0 {
				currPosition += 100
			}
		case "R":
			currPosition += position
			if currPosition >= 100 {
				currPosition -= 100
			}
		}

		if currPosition == 0 {
			zeroCounter++
		}
	}

	log.Printf("a: %d", zeroCounter)
}
