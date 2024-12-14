package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operator func(o1 int, o2 int) int

func EvaluateEquation(operands []int, operators []Operator, result int) bool {
	if len(operands) == 1 {
		return operands[0] == result
	}
	if operands[0] > result {
		return false
	}

	for _, operator := range operators {
		if EvaluateEquation(append([]int{operator(operands[0], operands[1])}, operands[2:]...), operators, result) {
			return true
		}
	}

	return false
}

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	equations := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		equation := strings.Split(line, ":")

		v, _ := strconv.Atoi(equation[0])
		var numbers []int

		for _, num := range strings.Split(strings.TrimSpace(equation[1]), " ") {
			v, _ := strconv.Atoi(num)
			numbers = append(numbers, v)
		}
		equations[v] = numbers
	}

	add := func(o1 int, o2 int) int {
		return o1 + o2
	}
	mul := func(o1 int, o2 int) int {
		return o1 * o2
	}
	concat := func(o1 int, o2 int) int {
		// Approach 1: strings
		// c := strconv.Itoa(o1) + strconv.Itoa(o2)
		// result, _ := strconv.Atoi(c)
		// Approach 2: shifts
		digits := int(math.Log10(float64(o2))) + 1
		shifted := o1 * int(math.Pow10(digits))
		return shifted + o2
	}

	sum := 0
	for result, operands := range equations {
		if EvaluateEquation(operands, []Operator{add, mul}, result) {
			sum += result
		}
	}

	fmt.Println("Puzzle 1:", sum)

	sum = 0
	for result, operands := range equations {
		if EvaluateEquation(operands, []Operator{add, mul, concat}, result) {
			sum += result
		}
	}
	fmt.Println("Puzzle 2:", sum)
}
