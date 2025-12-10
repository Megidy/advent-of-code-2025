package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Left  int
	Right int
}

// Could be imporved by using heap as a left start of range?
func main() {
	data, _ := os.ReadFile("input.txt")
	var totalCounter int

	re := regexp.MustCompile(`(?m)(?:\r?\n){2,}`)
	blocks := re.Split(strings.TrimSpace(string(data)), -1)

	rangesBlock := strings.Split(strings.TrimSpace(blocks[0]), "\n")
	numsBlock := strings.Split(strings.TrimSpace(blocks[1]), "\n")

	var ranges = make([]Range, 0, len(rangesBlock))

	for _, r := range rangesBlock {

		r = strings.TrimSpace(r)

		left, _ := strconv.Atoi(strings.Split(r, "-")[0])
		right, _ := strconv.Atoi(strings.Split(r, "-")[1])
		rangeOfBlock := Range{
			Left:  left,
			Right: right,
		}

		ranges = append(ranges, rangeOfBlock)

	}

	for _, numStr := range numsBlock {

		num, _ := strconv.Atoi(strings.TrimSpace(numStr))
		for _, r := range ranges {
			if r.Left <= num && num <= r.Right {
				totalCounter++
				break
			}
		}
	}

	log.Println(totalCounter)
}
