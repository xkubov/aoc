package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CityMap []string

type Vector struct {
	x int
	y int
}

func Combinations(vectors []Vector) (combinations [][]Vector) {
	for i, v := range vectors {
		for j := i + 1; j < len(vectors); j++ {
			combinations = append(combinations, []Vector{v, vectors[j]})
		}
	}
	return
}

func (v *Vector) Hash() string {
	return strconv.Itoa(v.x) + "," + strconv.Itoa(v.y)
}

func (cm *CityMap) GetAntenasPosition() map[string][]Vector {
	positions := make(map[string][]Vector)

	for i, line := range *cm {
		for j, char := range line {
			if string(char) != "." {
				positions[string(char)] = append(positions[string(char)], Vector{i, j})
			}
		}
	}

	return positions
}

func (cm *CityMap) ComputeAntinodes(point_a Vector, point_b Vector, harmonic bool) (antinodes []Vector) {
	if harmonic {
		antinodes = append(antinodes, point_a, point_b)
	}

	if point_a.x > point_b.x {
		point_a, point_b = point_b, point_a
	}

	diff_x := point_b.x - point_a.x
	diff_y := point_b.y - point_a.y

	for !cm.IsOutside(point_a) || !cm.IsOutside(point_b) {
		point_a = Vector{point_a.x - diff_x, point_a.y - diff_y}
		point_b = Vector{point_b.x + diff_x, point_b.y + diff_y}
		if !cm.IsOutside(point_a) {
			antinodes = append(antinodes, point_a)
		}
		if !cm.IsOutside(point_b) {
			antinodes = append(antinodes, point_b)
		}
		if !harmonic {
			break
		}
	}

	return
}

func (cm *CityMap) ComputeAllAntinodes(harmonic bool) (antinodes []Vector) {
	seen_antinodes := make(map[string]bool)
	for _, positions := range cm.GetAntenasPosition() {
		for _, points := range Combinations(positions) {
			for _, antinode := range cm.ComputeAntinodes(points[0], points[1], harmonic) {
				_, seen := seen_antinodes[antinode.Hash()]
				if !seen {
					seen_antinodes[antinode.Hash()] = true
					antinodes = append(antinodes, antinode)
				}
			}
		}
	}

	return
}

func (cm *CityMap) IsOutside(v Vector) bool {
	return v.x < 0 || v.x >= len(*cm) || v.y < 0 || v.y >= len((*cm)[0])
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	var city_map CityMap

	for scanner.Scan() {
		city_map = append(city_map, strings.TrimSpace(scanner.Text()))
	}

	antinodes := city_map.ComputeAllAntinodes(false)
	fmt.Println("Puzzle 1:", len(antinodes))

	antinodes = city_map.ComputeAllAntinodes(true)
	fmt.Println("Puzzle 2:", len(antinodes))
}
