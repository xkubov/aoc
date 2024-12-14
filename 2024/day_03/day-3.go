package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SumMuls(input string) int {
	r, _ := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)")

	sum := 0
	for _, v := range r.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])
		sum += x * y
	}
	return sum
}

func FilterDoBlocks(input string) string {
	var doblocks string

	r, _ := regexp.Compile("(?:^|do\\(\\))(.+?)(?:don't\\(\\)|$)")

	for _, v := range r.FindAllStringSubmatch(input, -1) {
		doblocks += v[1]
	}
	return doblocks
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input string
	for scanner.Scan() {
		input += strings.TrimSuffix(scanner.Text(), "\n")
	}

	fmt.Println("Puzzle 1:", SumMuls(input))
	fmt.Println("Puzzle 2:", SumMuls(FilterDoBlocks(input)))

	file.Close()
}
