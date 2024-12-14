package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	prints_before := make(map[int]([]int))

	// Scan rules
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		rule := strings.Split(line, "|")
		v1, _ := strconv.Atoi(rule[0])
		v2, _ := strconv.Atoi(rule[1])

		prints_before[v1] = append(prints_before[v1], v2)
	}

	var updates [][]int

	// Scan updates
	for scanner.Scan() {
		var update []int
		for _, v := range strings.Split(scanner.Text(), ",") {
			page, _ := strconv.Atoi(v)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	valid_sum := 0
	invalid_sum := 0

	for _, update := range updates {
		is_valid := true
		seen := make(map[int]bool)
		for _, page := range update {
			for _, required_after := range prints_before[page] {
				if _, ok := seen[required_after]; ok {
					is_valid = false
					break
				}
			}
			seen[page] = true
		}
		if is_valid {
			valid_sum += update[len(update)/2]
		} else {
			sort.Slice(update, func(i, j int) bool {
				return !slices.Contains(prints_before[update[j]], update[i])
			})
			invalid_sum += update[len(update)/2]
		}
	}
	fmt.Println("Puzzle 1:", valid_sum)
	fmt.Println("Puzzle 2:", invalid_sum)

	file.Close()
}
