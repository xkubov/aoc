package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func IsUnsafe(line []int) bool {
	last_increment := 0
	last_element := line[0]

	for _, v := range line[1:] {
		increment := v - last_element
		if last_increment*increment < 0 || !slices.Contains([]int{-1, -2, -3, 1, 2, 3}, increment) {
			return true
		}
		last_increment = increment
		last_element = v
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []([]int)

	for scanner.Scan() {
		var line []int
		for _, v := range strings.Split(scanner.Text(), " ") {
			i, _ := strconv.Atoi(v)
			line = append(line, i)
		}
		lines = append(lines, line)
	}

	unsafe := 0

	for _, line := range lines {
		if IsUnsafe(line) {
			unsafe++
		}
	}

	fmt.Println("Puzzle 1:")
	fmt.Println("    > Safe  :", len(lines)-unsafe)
	fmt.Println("    > Unsafe:", unsafe)

	unsafe = 0
	for _, line := range lines {
		if IsUnsafe(line) {
			safe := false
			for i := range line {
				if !IsUnsafe(RemoveIndex(line, i)) {
					safe = true
					break
				}
			}
			if !safe {
				unsafe++
			}
		}
	}

	fmt.Println("Puzzle 2:")
	fmt.Println("    > Safe  :", len(lines)-unsafe)
	fmt.Println("    > Unsafe:", unsafe)

	file.Close()
}
