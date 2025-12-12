package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type SquareTilePos struct {
	X int
	Y int
}

// сouldn't provide a more efficient solution; this one is O(n^2)
func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := make([][]rune, len(lines))
	var squareTilePositions = make([]SquareTilePos, 0)
	for i, line := range lines {
		grid[i] = []rune(line)
		split := strings.Split(string(line), ",")
		x, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(split[1]))
		squareTilePositions = append(squareTilePositions, SquareTilePos{
			X: x,
			Y: y,
		})
	}

	var maxRectangleArea int

	for i := 0; i < len(squareTilePositions)-1; i++ {
		for j := i; j < len(squareTilePositions)-1; j++ {
			width := math.Abs(float64(squareTilePositions[i].X)-float64(squareTilePositions[j].X)) + 1
			height := math.Abs(float64(squareTilePositions[i].Y)-float64(squareTilePositions[j].Y)) + 1
			area := int(width * height)

			if area > maxRectangleArea {
				maxRectangleArea = area
			}
		}
	}

	log.Println(maxRectangleArea)
}
