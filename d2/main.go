package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), ",")

	var totalCounter int

	for _, line := range lines {
		ids := strings.Split(line, "-")
		id1Str := ids[0]
		id2Str := ids[1]

		id1, _ := strconv.Atoi(id1Str)
		id2, _ := strconv.Atoi(id2Str)

		for id1 <= id2 {
			currId := id1

			currIdStr := strconv.Itoa(currId)
			if len(currIdStr)%2 == 0 {
				pow := int(math.Pow(10, float64(len(currIdStr)/2)))

				leftPart := currId / pow
				rightPart := currId % pow

				if leftPart == rightPart {
					totalCounter += currId
				}

			}

			id1++

		}
	}
	log.Println(totalCounter)

}
