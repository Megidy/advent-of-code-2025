package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var splitCounter int

	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	log.Println(lines)

	lines[1][len(lines[0])/2] = '|'

	for i := 1; i < len(lines)-1; i++ {
		for j := 0; j < len(lines)-1; j++ {
			previousValue := lines[i-1][j]

			if previousValue == '|' {
				currValue := lines[i][j]
				if currValue == '^' {
					lines[i][j-1] = '|'
					lines[i][j+1] = '|'
					splitCounter++
				} else {
					lines[i][j] = '|'
				}
			}
		}
	}

	log.Println(splitCounter)
	for _, row := range lines {
		log.Println(string(row))
	}
}
