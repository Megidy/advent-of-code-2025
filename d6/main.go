package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows [][]int
	var ops []string

	for i := 0; i < 4; i++ {
		scanner.Scan()

		lines := strings.Fields(scanner.Text())
		var row []int

		for _, line := range lines {
			n, _ := strconv.Atoi(line)

			row = append(row, n)
		}

		rows = append(rows, row)
	}

	scanner.Scan()

	ops = strings.Fields(scanner.Text())

	var totalCounter int

	rowLen := len(rows[0])

	for i := 0; i < rowLen; i++ {
		first := rows[0][i]
		second := rows[1][i]
		third := rows[2][i]
		fourth := rows[3][i]

		op := ops[i]
		var result int

		switch op {
		case "+":
			result = first + second + third + fourth
		case "*":
			result = first * second * third * fourth
		}
		totalCounter += result
	}

	log.Println(totalCounter)

}
