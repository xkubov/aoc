package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CountXMASHorizontal(lines []string) (sum int) {
	r1, _ := regexp.Compile("XMAS")
	r2, _ := regexp.Compile("SAMX")
	for _, line := range lines {
		sum += len(r1.FindAllString(line, -1)) + len(r2.FindAllString(line, -1))
	}
	return
}

func CountCrossMas(lines []string, left_d []string) (sum int) {
	r1, _ := regexp.Compile("MAS")
	r2, _ := regexp.Compile("SAM")

	for i := range left_d {
		m1_1 := r1.FindAllStringIndex(left_d[i], -1)
		m1_2 := r2.FindAllStringIndex(left_d[i], -1)

		for _, m := range append(m1_1, m1_2...) {
			pos_x := i
			pos_y := i%len(lines[0]) + 1
			if i >= len(lines[0]) {
				pos_x = len(lines[0]) - 1
			} else {
				pos_y = 0
			}
			pos_x -= m[0]
			pos_y += m[0]
			var right_str string
			l := m[1] - m[0]
			for i := 0; i < l; i++ {
				right_str += string(lines[pos_x-i][pos_y-i+l-1])
			}
			if r1.MatchString(right_str) || r2.MatchString(right_str) {
				sum++
			}
		}
	}
	return
}

func TransposeLines(lines []string) (transposed []string) {
	for i := 0; i < len(lines[0]); i++ {
		var line string
		for j := 0; j < len(lines); j++ {
			line += string(lines[j][i])
		}
		transposed = append(transposed, line)
	}
	return
}

func TransformRightDiagonal(lines []string) (transposed []string) {
	for i := 0; i < len(lines); i++ {
		var line string
		for j := 0; j <= i; j++ {
			line += string(lines[len(lines)-i-1+j][j])
		}
		transposed = append(transposed, line)
	}
	for i := 1; i < len(lines); i++ {
		var line string
		for j := 0; j <= len(lines)-i-1; j++ {
			line += string(lines[j][j+i])
		}
		transposed = append(transposed, line)
	}
	return
}

func TransformLeftDiagonal(lines []string) (transposed []string) {
	for i := 0; i < len(lines); i++ {
		var line string
		for j := 0; j <= i; j++ {
			line += string(lines[i-j][j])
		}
		transposed = append(transposed, line)
	}
	for i := 1; i < len(lines); i++ {
		var line string
		for j := 0; j <= len(lines)-i-1; j++ {
			line += string(lines[len(lines)-j-1][i+j])
		}
		transposed = append(transposed, line)
	}
	return
}

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\n"))
	}

	sum := CountXMASHorizontal(lines)
	sum += CountXMASHorizontal(TransposeLines(lines))
	sum += CountXMASHorizontal(TransformLeftDiagonal(lines))
	sum += CountXMASHorizontal(TransformRightDiagonal(lines))

	fmt.Println("Puzzle 1:")
	fmt.Println("    >", sum)

	sum = CountCrossMas(lines, TransformLeftDiagonal(lines))
	fmt.Println("Puzzle 2:")
	fmt.Println("    >", sum)

	file.Close()
}
