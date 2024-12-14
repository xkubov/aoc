package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")

	var nums_a []int
	var nums_b []int

	for {
		var a int
		var b int
		_, err := fmt.Fscanf(file, "%d %d\n", &a, &b)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}

		nums_a = append(nums_a, a)
		nums_b = append(nums_b, b)
	}

	sort.Ints(nums_a)
	sort.Ints(nums_b)

	sum := 0
	for i, v := range nums_a {
		sum += int(math.Abs(float64(v - nums_b[i])))
	}
	fmt.Println("puzzle 1:", sum)

	occurrences := make(map[int]int)
	for _, num := range nums_b {
		occurrences[num] = occurrences[num] + 1
	}

	sum = 0
	for i, v := range nums_a {
		sum += occurrences[v] * nums_a[i]
	}
	fmt.Println("puzzle 2:", sum)

	file.Close()
}
