package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	number int
}

func (s Stone) Blink() []Stone {
	if s.number == 0 {
		return []Stone{{number: 1}}
	}

	digits := int(math.Log10(float64(s.number)) + 1)

	if digits%2 == 0 {
		mask := int(math.Pow10(digits / 2))
		left := s.number / mask
		right := s.number % mask

		return []Stone{{number: left}, {number: right}}
	}

	return []Stone{{number: s.number * 2024}}
}

func (s Stone) BlinkN(n int, cache map[string]int) int {
	if n == 0 {
		return 1
	}
	h := Hash(s, n)

	if v, ok := cache[h]; ok {
		return v
	}

	sum := 0
	for _, ns := range s.Blink() {
		sum += ns.BlinkN(n-1, cache)
	}

	cache[h] = sum

	return sum
}

func Hash(stone Stone, num int) string {
	return strconv.Itoa(stone.number) + "," + strconv.Itoa(num)
}

func BlinkN(stones []Stone, n int) int {
	cache := make(map[string]int)

	sum := 0
	for _, stone := range stones {
		sum += stone.BlinkN(n, cache)
	}
	return sum
}

func GetStones(line string) []Stone {
	var stones []Stone
	split := strings.Split(strings.TrimSpace(line), " ")
	for _, val := range split {
		num, _ := strconv.Atoi(val)
		stones = append(stones, Stone{number: num})
	}
	return stones
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()

	stones := GetStones(line)

	fmt.Println("Puzzle 1:", BlinkN(stones, 25))
	fmt.Println("Puzzle 2:", BlinkN(stones, 75))

	f.Close()
}
